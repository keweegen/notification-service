package entity

import (
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/messagetemplate"
	"github.com/volatiletech/sqlboiler/v4/types"
	"time"
)

const (
	MessageStatusNew     = "new"
	MessageStatusSending = "sending"
	MessageStatusSent    = "sent"
	MessageStatusFailed  = "failed"
)

type Message struct {
	ID              string
	UserID          int64
	Channel         channel.Channel
	MessageTemplate messagetemplate.MessageTemplate
	Timestamp       int64
	ExternalID      int64
	Params          types.JSON
}

type Messages []*Message

type MessageStatus struct {
	ID          int64
	MessageID   string
	Status      string
	Description string
	CreatedAt   time.Time
}
