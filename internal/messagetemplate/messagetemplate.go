package messagetemplate

import "strings"

//go:generate stringer -type=MessageTemplate
type MessageTemplate int

const (
    Receipt MessageTemplate = iota + 1
)

var MessageTemplates = []MessageTemplate{Receipt}

func (i MessageTemplate) IsValid() bool {
    for _, mt := range MessageTemplates {
        if mt == i {
            return true
        }
    }

    return false
}

func GetMessageTemplateTypeFromString(s string) (MessageTemplate, bool) {
    switch strings.ToLower(s) {
    case strings.ToLower(Receipt.String()):
        return Receipt, true
    default:
        return 0, false
    }
}
