package repository

import (
    "database/sql"
    "github.com/go-redis/redis/v9"
)

type Store struct {
    Message Message
    User    User
}

func NewStore(db *sql.DB, mb redis.UniversalClient) *Store {
    return &Store{
        Message: new(messageRepository).init(db, mb),
        User:    new(userRepository).init(db),
    }
}
