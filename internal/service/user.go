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

func (u *User) UpdateNotificationChannels(ctx context.Context, channels entity.UserChannels) error {
    return u.repo.CreateOrUpdateChannels(ctx, channels)
}
