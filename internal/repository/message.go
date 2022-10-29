package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/internal/messagetemplate"
	"github.com/keweegen/notification/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type Message interface {
	Create(ctx context.Context, message *entity.Message) error
	CreateStatus(ctx context.Context, messageID, status, description string) error
	Find(ctx context.Context, messageID string) (*entity.Message, error)
	FindLastStatus(ctx context.Context, messageID string) (*entity.MessageStatus, error)
	Exists(ctx context.Context, messageID string) (bool, error)
	CheckForDuplicates(ctx context.Context, message *entity.Message) (string, error)
	Publish(ctx context.Context, key string, messageID string) error
	Subscribe(ctx context.Context, keys ...string) *MessageSubscription
}

type messageRepository struct {
	db *sql.DB
	mb redis.UniversalClient
}

func (r *messageRepository) init(db *sql.DB, mb redis.UniversalClient) Message {
	r.db = db
	r.mb = mb
	return r
}

func (r *messageRepository) Create(ctx context.Context, message *entity.Message) error {
	model := r.entityMessageToSqlboiler(message)
	if err := model.Insert(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}
	return nil
}

func (r *messageRepository) CreateStatus(ctx context.Context, messageID, status, description string) error {
	model := new(models.MessageStatus)
	model.MessageID = messageID
	model.Status = status
	model.Description = description

	if err := model.Insert(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to create message status: %w", err)
	}

	return nil
}

func (r *messageRepository) Find(ctx context.Context, messageID string) (*entity.Message, error) {
	model, err := models.FindMessage(ctx, r.db, messageID)
	if err != nil {
		return nil, fmt.Errorf("failed to find message: %w", err)
	}
	return r.sqlboilerToEntityMessage(model), nil
}

func (r *messageRepository) FindLastStatus(ctx context.Context, messageID string) (*entity.MessageStatus, error) {
	model, err := models.MessageStatuses(
		models.MessageStatusWhere.MessageID.EQ(messageID),
		qm.OrderBy(fmt.Sprintf("%s DESC", models.MessageStatusColumns.CreatedAt))).
		One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to find last status message: %w", err)
	}
	return r.sqlboilerToEntityMessageStatus(model), nil
}

func (r *messageRepository) Exists(ctx context.Context, messageID string) (bool, error) {
	exist, err := models.MessageExists(ctx, r.db, messageID)
	if err != nil {
		return false, fmt.Errorf("failed to check exists message: %w", err)
	}
	return exist, nil
}

func (r *messageRepository) CheckForDuplicates(ctx context.Context, message *entity.Message) (string, error) {
	offset := time.Minute * 5

	model, err := models.Messages(
		models.MessageWhere.UserID.EQ(message.UserID),
		models.MessageWhere.ExternalID.EQ(message.ExternalID),
		models.MessageWhere.Channel.EQ(int16(message.Channel)),
		models.MessageWhere.Template.EQ(int16(message.MessageTemplate)),
		models.MessageWhere.Timestamp.GTE(time.UnixMilli(message.Timestamp).Add(-offset)),
		models.MessageWhere.Timestamp.LTE(time.UnixMilli(message.Timestamp).Add(offset))).
		One(ctx, r.db)
	if err != nil {
		return "", fmt.Errorf("failed to check for duplicates message: %w", err)
	}
	return model.ID, nil
}

func (r *messageRepository) Publish(ctx context.Context, key string, messageID string) error {
	return r.mb.Publish(ctx, key, messageID).Err()
}

func (r *messageRepository) Subscribe(ctx context.Context, keys ...string) *MessageSubscription {
	return &MessageSubscription{pubSub: r.mb.Subscribe(ctx, keys...)}
}

func (r *messageRepository) entityMessageToSqlboiler(data *entity.Message) *models.Message {
	return &models.Message{
		ID:         data.ID,
		UserID:     data.UserID,
		ExternalID: data.ExternalID,
		Channel:    int16(data.Channel),
		Template:   int16(data.MessageTemplate),
		Timestamp:  time.UnixMilli(data.Timestamp),
		Params:     data.Params,
	}
}

func (r *messageRepository) sqlboilerToEntityMessage(data *models.Message) *entity.Message {
	return &entity.Message{
		ID:              data.ID,
		UserID:          data.UserID,
		ExternalID:      data.ExternalID,
		Channel:         channel.Channel(data.Channel),
		MessageTemplate: messagetemplate.MessageTemplate(data.Template),
		Timestamp:       data.Timestamp.UnixMilli(),
		Params:          data.Params,
	}
}

func (r *messageRepository) entityMessageStatusToSqlboiler(data *entity.MessageStatus) *models.MessageStatus {
	return &models.MessageStatus{
		ID:          data.ID,
		MessageID:   data.MessageID,
		Status:      data.Status,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
	}
}

func (r *messageRepository) sqlboilerToEntityMessageStatus(data *models.MessageStatus) *entity.MessageStatus {
	return &entity.MessageStatus{
		ID:          data.ID,
		MessageID:   data.MessageID,
		Status:      data.Status,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
	}
}

type MessageSubscription struct {
	pubSub *redis.PubSub
}

func (ms *MessageSubscription) Receive(ctx context.Context) (string, string, error) {
	if ms.pubSub == nil {
		return "", "", errors.New("pubsub instance not set")
	}

	msg, err := ms.pubSub.ReceiveMessage(ctx)
	if err != nil {
		return "", "", err
	}

	return msg.Channel, msg.Payload, nil
}

func (ms *MessageSubscription) Close() error {
	if ms.pubSub == nil {
		return nil
	}
	return ms.pubSub.Close()
}
