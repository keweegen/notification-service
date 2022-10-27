package email

import (
    "bytes"
    "fmt"
    "github.com/google/uuid"
    "net/smtp"
    "time"
)

type client struct {
    host string
    port uint
    from string
    auth smtp.Auth
}

func (c *client) init(host string, port uint, from, username, password string) *client {
    c.host = host
    c.port = port
    c.from = from
    c.auth = smtp.PlainAuth("", username, password, c.host)
    return c
}

func (c *client) do(to, content string) error {
    var body bytes.Buffer
    body.WriteString(fmt.Sprintf("%s\n\n%s", c.makeHeaders(to), content))

    if err := smtp.SendMail(c.smtpAddress(), c.auth, c.from, []string{to}, body.Bytes()); err != nil {
        return fmt.Errorf("failed send email: %w", err)
    }

    return nil
}

func (c *client) makeHeaders(to string) string {
    mid := fmt.Sprintf("<%s:%s>", uuid.NewString(), c.from)
    subject := "Notification Service"
    mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

    return fmt.Sprintf("Message-ID: %s\nDate: %s\nFrom: %s\nTo: %s\nSubject: %s\n%s",
        mid,
        time.Now().Format(time.RFC1123Z),
        c.from,
        to,
        subject,
        mime)
}

func (c *client) smtpAddress() string {
    return fmt.Sprintf("%s:%d", c.host, c.port)
}
