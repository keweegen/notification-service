package messagetemplate

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
)

func TestGetMessageTemplateTypeFromString(t *testing.T) {
    type testCase struct {
        name  string
        input string
        mt    MessageTemplate
        isOK  bool
    }

    cases := []testCase{
        {name: "unknown type", input: "unknown", mt: 0, isOK: false},
    }

    for _, mt := range MessageTemplates {
        cases = append(cases,
            testCase{
                name:  mt.String(),
                input: mt.String(),
                mt:    mt,
                isOK:  true,
            },
            testCase{
                name:  fmt.Sprintf("%s lowercase", mt.String()),
                input: strings.ToLower(mt.String()),
                mt:    mt,
                isOK:  true,
            },
        )
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            mt, ok := GetMessageTemplateTypeFromString(tc.input)
            assert.Equal(t, tc.mt, mt, "is equal message template type")
            assert.Equal(t, tc.isOK, ok, "is exists message template type")
            assert.Equal(t, tc.isOK, mt.IsValid(), "is valid message template type")
        })
    }
}
