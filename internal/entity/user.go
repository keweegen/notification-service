package entity

import "github.com/keweegen/notification/internal/channel"

type UserChannel struct {
	ID        int64
	UserID    int64
	Channel   channel.Channel
	Recipient string
	CanNotify bool
}

type UserChannels []*UserChannel
