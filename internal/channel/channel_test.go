package channel

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
)

func TestGetChannelTypeFromString(t *testing.T) {
    type testCase struct {
        name  string
        input string
        ch    Channel
        isOK  bool
    }

    cases := []testCase{
        {name: "unknown type", input: "unknown", ch: 0, isOK: false},
    }

    for _, ch := range Channels {
        cases = append(cases,
            testCase{
                name:  ch.String(),
                input: ch.String(),
                ch:    ch,
                isOK:  true,
            },
            testCase{
                name:  fmt.Sprintf("%s lowercase", ch.String()),
                input: strings.ToLower(ch.String()),
                ch:    ch,
                isOK:  true,
            },
        )
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            ch, ok := GetChannelTypeFromString(tc.input)
            assert.Equal(t, tc.ch, ch, "is equal channel type")
            assert.Equal(t, tc.isOK, ok, "is exists channel type")
            assert.Equal(t, tc.isOK, ch.IsValid(), "is valid channel type")
        })
    }
}
