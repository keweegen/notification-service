package service

import (
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/repository"
	"github.com/keweegen/notification/logger"
)

type Store struct {
	Message        *Message
	MessageChecker *MessageChecker
	User           *User
}

func NewStore(l logger.Logger, repo *repository.Store, channels *channel.Store) *Store {
	m := NewMessage(l, repo, channels)

	return &Store{
		Message:        m,
		MessageChecker: NewMessageChecker(l, repo, m),
		User:           NewUser(repo.User),
	}
}
