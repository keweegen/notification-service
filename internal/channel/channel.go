package channel

import "strings"

//go:generate stringer -type=Channel
type Channel int

const (
    Telegram Channel = iota + 1
    Email
)

var Channels = []Channel{Telegram, Email}

func (i Channel) IsValid() bool {
    for _, c := range Channels {
        if c == i {
            return true
        }
    }

    return false
}

func GetChannelTypeFromString(s string) (Channel, bool) {
    switch strings.ToLower(s) {
    case strings.ToLower(Telegram.String()):
        return Telegram, true
    case strings.ToLower(Email.String()):
        return Email, true
    default:
        return 0, false
    }
}
