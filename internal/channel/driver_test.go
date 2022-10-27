package channel

import (
    "github.com/keweegen/notification/config"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestStore_Get(t *testing.T) {
    store := NewStore(config.NotificationChannels{})
    driverTelegram, _ := store.Get(Telegram)
    driverEmail, _ := store.Get(Email)

    cases := []struct {
        name           string
        inputChannel   Channel
        expectedDriver Driver
        expectedError  error
    }{
        {
            name:           "unknown driver",
            inputChannel:   Channel(0),
            expectedDriver: nil,
            expectedError:  DriverNotFoundErr,
        },
        {
            name:           "telegram driver",
            inputChannel:   Telegram,
            expectedDriver: driverTelegram,
            expectedError:  nil,
        },
        {
            name:           "email driver",
            inputChannel:   Email,
            expectedDriver: driverEmail,
            expectedError:  nil,
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            driver, err := store.Get(tc.inputChannel)
            assert.Equal(t, tc.expectedError, err)
            assert.Equal(t, tc.expectedDriver, driver)
        })
    }
}
