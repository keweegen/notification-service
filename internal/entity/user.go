package entity

import "github.com/keweegen/notification/internal/channel"

type UserChannel struct {
    UserID    int64
    Channel   channel.Channel
    Recipient string
    CanNotify bool
}

type UserChannels []*UserChannel
