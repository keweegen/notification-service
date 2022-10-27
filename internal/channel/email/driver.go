package email

import "github.com/keweegen/notification/config"

type Driver struct {
    client *client
}

func New(cfg config.Email) *Driver {
    return &Driver{
        client: new(client).init(cfg.Host, cfg.Port, cfg.From, cfg.Username, cfg.Password),
    }
}

func (d *Driver) Send(receiver, message string) error {
    return d.client.do(receiver, message)
}
