package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/internal/messagetemplate"
	"github.com/keweegen/notification/internal/repository"
	"github.com/keweegen/notification/logger"
	"github.com/volatiletech/sqlboiler/v4/types"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	retryTimeout = 5 * time.Second
)

var (
	MessageNotFoundErr        = errors.New("message not found")
	InvalidMessageIDErr       = errors.New("messageId: invalid")
	InvalidChannelErr         = errors.New("messageId: invalid channel")
	InvalidMessageTemplateErr = errors.New("messageId: invalid message template")
	InvalidUserErr            = errors.New("messageId: invalid user")
	InvalidTimestamp          = errors.New("messageId: invalid timestamp")

	minTime = time.Now().AddDate(-3, 0, 0)
	maxTime = time.Now().AddDate(1, 0, 0)
)

type Message struct {
	logger       logger.Logger
	repoStore    *repository.Store
	channelStore *channel.Store

	chQueueChannels map[channel.Channel]chan string
	mx              *sync.Mutex
}

func NewMessage(l logger.Logger, repo *repository.Store, channelStore *channel.Store) *Message {
	channels := make(map[channel.Channel]chan string)

	for _, ch := range channel.Channels {
		channels[ch] = make(chan string)
	}

	return &Message{
		logger:          l.With("service", "message"),
		repoStore:       repo,
		channelStore:    channelStore,
		chQueueChannels: channels,
		mx:              new(sync.Mutex),
	}
}

func (m *Message) GenerateID(
	channel channel.Channel,
	messageTemplate messagetemplate.MessageTemplate,
	userID int64,
	timestamp int64,
	externalID int64,
) string {
	return fmt.Sprintf("NS-%03d-%03d-%019d-%s-%019d",
		channel,
		messageTemplate,
		userID,
		m.padRightSide(timestamp),
		externalID)
}

func (m *Message) ValidateID(ctx context.Context, id string) error {
	_, err := m.parseID(ctx, id)
	return err
}

func (m *Message) parseID(ctx context.Context, id string) (*entity.Message, error) {
	if id == "" {
		return nil, InvalidMessageIDErr
	}

	items := strings.Split(id, "-")
	if len(items) != 6 {
		return nil, InvalidMessageIDErr
	}

	channelID, _ := strconv.Atoi(items[1])
	messageTemplateID, _ := strconv.Atoi(items[2])
	userID, _ := strconv.Atoi(items[3])
	timestamp, _ := strconv.Atoi(items[4])
	externalID, _ := strconv.Atoi(items[5])

	ch := channel.Channel(channelID)
	if !ch.IsValid() {
		return nil, InvalidChannelErr
	}

	mt := messagetemplate.MessageTemplate(messageTemplateID)
	if !mt.IsValid() {
		return nil, InvalidMessageTemplateErr
	}

	userExists, err := m.repoStore.User.Exists(ctx, int64(userID))
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, InvalidUserErr
	}

	if t := time.UnixMilli(int64(timestamp)); t.Before(minTime) || t.After(maxTime) {
		return nil, InvalidTimestamp
	}

	return &entity.Message{
		ID:              id,
		Channel:         ch,
		UserID:          int64(userID),
		MessageTemplate: mt,
		Timestamp:       int64(timestamp),
		ExternalID:      int64(externalID),
	}, nil
}

func (m *Message) GetStatus(ctx context.Context, id string) (*entity.MessageStatus, error) {
	if err := m.ValidateID(ctx, id); err != nil {
		return nil, fmt.Errorf("get status: %w", err)
	}

	messageStatus, err := m.repoStore.Message.FindLastStatus(ctx, id)
	if err == nil {
		return messageStatus, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, MessageNotFoundErr
	}

	return nil, err
}

func (m *Message) Send(ctx context.Context, id string, params types.JSON) (string, error) {
	message, err := m.parseID(ctx, id)
	if err != nil {
		return "", err
	}

	message.Params = params

	messageId, err := m.repoStore.Message.CheckForDuplicates(ctx, message)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	if messageId != "" {
		return messageId, nil
	}

	if err = m.repoStore.Message.Create(ctx, message); err != nil {
		return "", err
	}
	if err = m.repoStore.Message.CreateStatus(ctx, message.ID, entity.MessageStatusNew, ""); err != nil {
		return "", err
	}

	return message.ID, m.repoStore.Message.Publish(ctx, m.pubSubKey(message.Channel), message.ID)
}

func (m *Message) HandleMessages(ctx context.Context, quit <-chan struct{}) {
	subscriptionKeys := make([]string, 0, len(channel.Channels))

	for _, ch := range channel.Channels {
		subscriptionKeys = append(subscriptionKeys, m.pubSubKey(ch))

		go m.handleChannelMessages(ctx, quit, ch)
	}

	subscription := m.repoStore.Message.Subscribe(ctx, subscriptionKeys...)

	go func() {
		for {
			select {
			case <-quit:
				m.onCloseSubscriptionReceive(subscription)
				return
			default:
				m.receiveFromSubscription(ctx, subscription)
			}
		}
	}()
}

func (m *Message) onCloseSubscriptionReceive(subscription *repository.MessageSubscription) {
	if err := subscription.Close(); err != nil {
		m.logger.Error("failed close subscription message broker connection",
			"error", err)
	}
}

func (m *Message) receiveFromSubscription(ctx context.Context, subscription *repository.MessageSubscription) {
	ch, messageID, err := subscription.Receive(ctx)
	if err != nil {
		m.logger.Error("receive message from message broker", "error", err)
		time.Sleep(retryTimeout)
		return
	}

	m.appendQueueChannelMessage(ch, messageID)
}

func (m *Message) appendQueueChannelMessage(channel, messageID string) {
	ch, err := m.getChannelNameFromPubSubKey(channel)
	if err != nil {
		m.logger.Error("get channel from PubSub key", "channel", channel, "error", err)
		return
	}

	m.mx.Lock()
	defer m.mx.Unlock()

	m.chQueueChannels[ch] <- messageID
}

func (m *Message) getChannelNameFromPubSubKey(key string) (channel.Channel, error) {
	items := strings.Split(key, "::")
	if len(items) <= 1 {
		return 0, InvalidChannelErr
	}

	intChannel, _ := strconv.Atoi(items[1])
	ch := channel.Channel(intChannel)
	if !ch.IsValid() {
		return 0, InvalidChannelErr
	}

	return ch, nil
}

func (m *Message) padRightSide(n int64) string {
	num := fmt.Sprintf("%d", n)
	repeatCount := 13 - len(num)

	if repeatCount <= 0 {
		return num
	}

	return num + strings.Repeat("0", repeatCount)
}

func (m *Message) handleChannelMessages(ctx context.Context, quit <-chan struct{}, channel channel.Channel) {
	for {
		select {
		case <-quit:
			return
		default:
			messageID := <-m.chQueueChannels[channel]
			l := m.logger.With("messageId", messageID, "channel", channel)

			message, err := m.repoStore.Message.Find(ctx, messageID)
			if err != nil {
				l.Error("find message", "error", err)
				m.makeStatus(ctx, message.ID, entity.MessageStatusFailed, err.Error())
				continue
			}

			if err = m.sendMessage(ctx, message); err != nil {
				l.Error("failed send message", "error", err)
			}
		}
	}
}

func (m *Message) sendMessage(ctx context.Context, message *entity.Message) error {
	m.logger.Debug("sending message")
	m.makeStatus(ctx, message.ID, entity.MessageStatusSending, "Sending a message")

	channelDriver, err := m.channelStore.Get(message.Channel)
	if err != nil {
		return fmt.Errorf("get channel driver by name: %w", err)
	}

	userChannelSettings, err := m.repoStore.User.FindByChannel(ctx, message.UserID, message.Channel)
	if err != nil {
		return fmt.Errorf("find user notification channel: %w", err)
	}
	if !userChannelSettings.CanNotify {
		return errors.New("CanNotify is false")
	}

	content, err := m.getContentFromTemplate(message)
	if err != nil {
		return fmt.Errorf("get content from template")
	}
	if err = channelDriver.Send(userChannelSettings.Recipient, content); err != nil {
		return fmt.Errorf("send message with channel driver: %w", err)
	}

	m.makeStatus(ctx, message.ID, entity.MessageStatusSent, "Message sent")
	m.logger.Debug("message sent", "channel", message.Channel, "content", content)

	return nil
}

func (m *Message) makeStatus(ctx context.Context, messageID, status, description string) {
	if err := m.repoStore.Message.CreateStatus(ctx, messageID, status, description); err != nil {
		m.logger.Error("create message status",
			"messageId", messageID,
			"status", status,
			"description", description,
			"error", err)
	}
}

func (m *Message) pubSubKey(channel channel.Channel) string {
	return fmt.Sprintf("ns::%d", channel)
}

func (m *Message) getContentFromTemplate(message *entity.Message) (string, error) {
	tmpl, err := messagetemplate.GetTemplate(message.MessageTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to get message template: %w", err)
	}

	if err = tmpl.SetParams(message.Params); err != nil {
		return "", fmt.Errorf("failed to set params message template: %s", err)
	}

	data, err := messagetemplate.Parse(tmpl, message.Channel)
	if err != nil {
		return "", fmt.Errorf("failed to parse message template: %w", err)
	}

	return data, nil
}
