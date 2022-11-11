package service

import (
	"context"
	"fmt"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/internal/repository"
	"github.com/keweegen/notification/logger"
	"time"
)

const (
	retryTimeoutChecker = 10 * time.Minute
)

type MessageChecker struct {
	logger         logger.Logger
	repo           *repository.Store
	messageService *Message
}

func NewMessageChecker(logger logger.Logger, repo *repository.Store, messageService *Message) *MessageChecker {
	return &MessageChecker{
		logger:         logger.With("service", "messageChecker"),
		repo:           repo,
		messageService: messageService,
	}
}

func (mc *MessageChecker) Do(ctx context.Context) {
	mc.check(ctx)

	for {
		select {
		case <-ctx.Done():
			mc.logger.Debug("Do: context done")
			return
		case <-time.After(retryTimeoutChecker):
			mc.check(ctx)
		}
	}
}

func (mc *MessageChecker) check(ctx context.Context) {
	mc.logger.Debug("get process messages")

	processMessages, err := mc.getProcessMessages(ctx)
	if err != nil {
		mc.logger.Error("failed to get process messages", "error", err)
		return
	}

	mc.resendMessages(ctx, processMessages)
}

func (mc *MessageChecker) resendMessages(ctx context.Context, messages entity.Messages) {
	for _, message := range messages {
		message := message

		go func() {
			if err := mc.messageService.sendMessage(ctx, message); err != nil {
				mc.logger.Error("failed to send message",
					"messageId", message.ID,
					"error", err)
			}
		}()
	}
}

func (mc *MessageChecker) getProcessMessages(ctx context.Context) (entity.Messages, error) {
	dateStart := time.Now().Add(-24 * time.Hour)
	dateEnd := time.Now().Add(-2 * time.Hour)

	messages, err := mc.repo.Message.FindProcessMessages(ctx, dateStart, dateEnd)
	if err != nil {
		return nil, fmt.Errorf("get process messages: %w", err)
	}

	return messages, nil
}
