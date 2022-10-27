package messagetemplate

import (
    "bytes"
    "errors"
    "fmt"
    "github.com/keweegen/notification/internal/channel"
    "github.com/volatiletech/sqlboiler/v4/types"
    "html/template"
)

var TemplateNotFoundErr = errors.New("template not found")

type Template interface {
    Name() string
    SetParams(data types.JSON) error
    EmailTemplate() *template.Template
    TelegramTemplate() *template.Template
}

var templates = map[MessageTemplate]Template{
    Receipt: new(ReceiptTemplate),
}

func GetTemplate(name MessageTemplate) (Template, error) {
    tmpl, ok := templates[name]
    if !ok {
        return nil, TemplateNotFoundErr
    }
    return tmpl, nil
}

func Parse(t Template, ch channel.Channel) (string, error) {
    tmpl, err := getChannelTemplateByName(t, ch)
    if err != nil {
        return "", err
    }

    var result bytes.Buffer
    if err = tmpl.Execute(&result, t); err != nil {
        return "", fmt.Errorf("execute: %w", err)
    }

    return result.String(), nil
}

func getChannelTemplateByName(t Template, ch channel.Channel) (*template.Template, error) {
    switch ch {
    case channel.Email:
        return t.EmailTemplate(), nil
    case channel.Telegram:
        return t.TelegramTemplate(), nil
    default:
        return nil, fmt.Errorf("unknown channel '%s'", ch.String())
    }
}
