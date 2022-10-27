package app

import (
    "context"
    "database/sql"
    "github.com/go-redis/redis/v9"
    "github.com/keweegen/notification/config"
    "github.com/keweegen/notification/db"
    "github.com/keweegen/notification/logger"
    "github.com/keweegen/notification/messagebroker"
    "github.com/pkg/errors"
)

type App struct {
    Config *config.Config
    Logger logger.Logger

    db *sql.DB
    mb redis.UniversalClient
}

func New(cfg *config.Config, logger logger.Logger) *App {
    app := new(App)
    app.Config = cfg
    app.Logger = logger
    return app
}

func (a *App) Connect() error {
    if err := a.OpenConnectDatabase(); err != nil {
        return err
    }
    if err := a.openConnectMessageBroker(); err != nil {
        return err
    }
    return nil
}

func (a *App) Close() error {
    if err := a.CloseConnectDatabase(); err != nil {
        return err
    }
    if err := a.closeConnectMessageBroker(); err != nil {
        return err
    }
    return nil
}

func (a *App) CurrentDatabase() *sql.DB {
    return a.db
}

func (a *App) CurrentMessageBroker() redis.UniversalClient {
    return a.mb
}

func (a *App) OpenConnectDatabase() (err error) {
    cfg := a.Config.Database

    a.db, err = db.NewConnect(cfg.Host, cfg.Name, cfg.User, cfg.Password, cfg.Port)
    return errors.Wrap(err, "open connect database")
}

func (a *App) openConnectMessageBroker() (err error) {
    cfg := a.Config.MessageBroker

    a.mb, err = messagebroker.NewConnect(context.Background(), cfg.Password, cfg.Addr)
    return errors.Wrap(err, "open connect message broker")
}

func (a *App) CloseConnectDatabase() error {
    if a.db == nil {
        return nil
    }
    return a.db.Close()
}

func (a *App) closeConnectMessageBroker() error {
    if a.mb == nil {
        return nil
    }
    return a.mb.Close()
}
