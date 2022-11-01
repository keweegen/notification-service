package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/models"
)

//go:generate mockgen -source=user.go -destination=./mock/user.go
type User interface {
	CreateOrUpdateChannels(ctx context.Context, channels entity.UserChannels) error
	FindByChannel(ctx context.Context, userID int64, channel channel.Channel) (*entity.UserChannel, error)
	Exists(ctx context.Context, userID int64) (bool, error)
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) init(db *sql.DB) User {
	r.db = db
	return r
}

func (r *userRepository) CreateOrUpdateChannels(ctx context.Context, channels entity.UserChannels) error {
	return nil
}

func (r *userRepository) FindByChannel(ctx context.Context, userID int64, channel channel.Channel) (*entity.UserChannel, error) {
	data, err := models.UserChannels(
		models.UserChannelWhere.UserID.EQ(userID),
		models.UserChannelWhere.Channel.EQ(int16(channel))).
		One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to find user channels by channel: %w", err)
	}

	return r.sqlboilerToEntity(data), nil
}

func (r *userRepository) Exists(ctx context.Context, userID int64) (bool, error) {
	exist, err := models.UserChannels(
		models.UserChannelWhere.UserID.EQ(userID)).
		Exists(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("failed to check exists user channels: %w", err)
	}

	return exist, nil
}

func (r *userRepository) sqlboilerToEntity(data *models.UserChannel) *entity.UserChannel {
	return &entity.UserChannel{
		UserID:    data.UserID,
		Channel:   channel.Channel(data.Channel),
		Recipient: data.Recipient,
		CanNotify: data.CanNotify,
	}
}
