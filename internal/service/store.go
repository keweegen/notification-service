package service

import (
    "github.com/keweegen/notification/internal/channel"
    "github.com/keweegen/notification/internal/repository"
    "github.com/keweegen/notification/logger"
)

type Store struct {
    Message *Message
    User    *User
}

func NewStore(l logger.Logger, repo *repository.Store, channels *channel.Store) *Store {
    return &Store{
        Message: NewMessage(l, repo, channels),
        User:    NewUser(repo.User),
    }
}
