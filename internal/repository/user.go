package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var UserChannelNotFound = errors.New("user channel not found")

//go:generate mockgen -source=user.go -destination=./mock/user.go
type User interface {
	CreateChannel(ctx context.Context, channel *entity.UserChannel) error
	FindChannel(ctx context.Context, channelID int64) (*entity.UserChannel, error)
	UpdateChannel(ctx context.Context, channel *entity.UserChannel) error
	DestroyChannel(ctx context.Context, channelID int64) error
	FindChannelsByUser(ctx context.Context, userID int64) (entity.UserChannels, error)
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

func (r *userRepository) CreateChannel(ctx context.Context, channel *entity.UserChannel) error {
	model := r.entityToSqlboiler(channel)
	if err := model.Insert(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to create user channel: %w", err)
	}
	return nil
}

func (r *userRepository) FindChannel(ctx context.Context, channelID int64) (*entity.UserChannel, error) {
	model, err := models.FindUserChannel(ctx, r.db, channelID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, UserChannelNotFound
		}
		return nil, fmt.Errorf("failed to find user channel: %w", err)
	}
	return r.sqlboilerToEntity(model), nil
}

func (r *userRepository) UpdateChannel(ctx context.Context, channel *entity.UserChannel) error {
	if _, err := r.entityToSqlboiler(channel).Update(ctx, r.db, boil.Infer()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UserChannelNotFound
		}
		return fmt.Errorf("failed to update user channel: %w", err)
	}
	return nil
}

func (r *userRepository) DestroyChannel(ctx context.Context, channelID int64) error {
	if _, err := models.UserChannels(models.UserChannelWhere.ID.EQ(channelID)).DeleteAll(ctx, r.db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UserChannelNotFound
		}
		return fmt.Errorf("failed to destroy channel: %w", err)
	}
	return nil
}

func (r *userRepository) FindChannelsByUser(ctx context.Context, userID int64) (entity.UserChannels, error) {
	userChannels, err := models.UserChannels(models.UserChannelWhere.UserID.EQ(userID)).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, UserChannelNotFound
		}
		return nil, fmt.Errorf("failed to find channels by user: %w", err)
	}
	return r.sqlboilerToEntities(userChannels), nil
}

func (r *userRepository) FindByChannel(ctx context.Context, userID int64, channel channel.Channel) (*entity.UserChannel, error) {
	data, err := models.UserChannels(
		models.UserChannelWhere.UserID.EQ(userID),
		models.UserChannelWhere.Channel.EQ(int16(channel))).
		One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, UserChannelNotFound
		}
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
		ID:        data.ID,
		UserID:    data.UserID,
		Channel:   channel.Channel(data.Channel),
		Recipient: data.Recipient,
		CanNotify: data.CanNotify,
	}
}

func (r *userRepository) sqlboilerToEntities(data models.UserChannelSlice) entity.UserChannels {
	channels := make(entity.UserChannels, 0, len(data))

	for _, item := range data {
		channels = append(channels, r.sqlboilerToEntity(item))
	}

	return channels
}

func (r *userRepository) entityToSqlboiler(data *entity.UserChannel) *models.UserChannel {
	return &models.UserChannel{
		ID:        data.ID,
		UserID:    data.UserID,
		Channel:   int16(data.Channel),
		Recipient: data.Recipient,
		CanNotify: data.CanNotify,
	}
}
