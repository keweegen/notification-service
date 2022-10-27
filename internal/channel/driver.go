package channel

import (
    "errors"
    "github.com/keweegen/notification/config"
    "github.com/keweegen/notification/internal/channel/email"
    "github.com/keweegen/notification/internal/channel/telegram"
)

var (
    DriverNotFoundErr = errors.New("channel driver not found")
)

type Driver interface {
    Send(receiver, message string) error
}

type Store struct {
    drivers map[Channel]Driver
}

func NewStore(cfg config.NotificationChannels) *Store {
    return &Store{drivers: map[Channel]Driver{
        Telegram: telegram.New(cfg.Telegram),
        Email:    email.New(cfg.Email),
    }}
}

func (s *Store) Get(channel Channel) (Driver, error) {
    driver, ok := s.drivers[channel]
    if !ok {
        return nil, DriverNotFoundErr
    }
    return driver, nil
}
