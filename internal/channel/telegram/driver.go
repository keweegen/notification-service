package telegram

import "github.com/keweegen/notification/config"

type Driver struct {
    client *client
}

func New(cfg config.Telegram) *Driver {
    return &Driver{client: new(client).init(cfg.Host, cfg.APIKey)}
}

func (d *Driver) Send(receiver, message string) error {
    return d.client.SendMessage(receiver, message, "HTML")
}
