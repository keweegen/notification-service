package messagetemplate

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/keweegen/notification/internal/channel"
	mock_messagetemplate "github.com/keweegen/notification/internal/messagetemplate/mock"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	type testCase struct {
		name             string
		templateType     MessageTemplate
		expectedTemplate Template
		expectedError    error
	}

	cases := []testCase{
		{
			name:             "unknown message template type",
			templateType:     MessageTemplate(0),
			expectedTemplate: nil,
			expectedError:    TemplateNotFoundErr,
		},
	}

	for _, mt := range MessageTemplates {
		tmpl, ok := templates[mt]
		if !ok {
			t.Fatalf("message template not found in templates, messageTemplateType=%s", mt)
		}

		cases = append(cases, testCase{
			name:             fmt.Sprintf("%s message template type", mt.String()),
			templateType:     mt,
			expectedTemplate: tmpl,
			expectedError:    nil,
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tmpl, err := GetTemplate(tc.templateType)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTemplate, tmpl)
		})
	}
}

func TestParse(t *testing.T) {
	type testCase struct {
		name                    string
		channel                 channel.Channel
		mockTemplate            *template.Template
		expectedTemplateContent string
		expectedError           error
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	tmpl := mock_messagetemplate.NewMockTemplate(controller)

	cases := []testCase{
		{
			name:                    "template for unknown",
			channel:                 channel.Channel(0),
			mockTemplate:            nil,
			expectedTemplateContent: "",
			expectedError:           errors.New("unknown channel 'Channel(0)'"),
		},
	}

	for _, ch := range channel.Channels {
		key := fmt.Sprintf("ns-test.%s.mock", ch.String())
		content := fmt.Sprintf("Test template for %s", ch.String())

		cases = append(cases, testCase{
			name:                    fmt.Sprintf("template for %s", ch.String()),
			channel:                 ch,
			mockTemplate:            template.Must(template.New(key).Parse(content)),
			expectedTemplateContent: content,
			expectedError:           nil,
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.channel {
			case channel.Mock, channel.Telegram:
				tmpl.EXPECT().TelegramTemplate().Return(tc.mockTemplate)
			case channel.Email:
				tmpl.EXPECT().EmailTemplate().Return(tc.mockTemplate)
			}

			chTemplate, err := Parse(tmpl, tc.channel)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTemplateContent, chTemplate)
		})
	}
}
