package service

import (
	"context"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/internal/repository"
)

type User struct {
	repo repository.User
}

func NewUser(repo repository.User) *User {
	return &User{repo: repo}
}

func (u *User) CreateNotificationChannel(ctx context.Context, userChannel *entity.UserChannel) error {
	userChannel.ID = 0
	if !userChannel.Channel.IsValid() {
		return InvalidChannelErr
	}

	return u.repo.CreateChannel(ctx, userChannel)
}

func (u *User) FindNotificationChannel(ctx context.Context, userChannelID int64) (*entity.UserChannel, error) {
	return u.repo.FindChannel(ctx, userChannelID)
}

func (u *User) UpdateNotificationChannel(ctx context.Context, userChannelID int64, recipient string, canNotify bool) error {
	model, err := u.FindNotificationChannel(ctx, userChannelID)
	if err != nil {
		return err
	}

	model.Recipient = recipient
	model.CanNotify = canNotify

	return u.repo.UpdateChannel(ctx, model)
}

func (u *User) DestroyNotificationChannel(ctx context.Context, userChannelID int64) error {
	return u.repo.DestroyChannel(ctx, userChannelID)
}

func (u *User) FindNotificationChannels(ctx context.Context, userID int64) (entity.UserChannels, error) {
	return u.repo.FindChannelsByUser(ctx, userID)
}
