package http

import (
	"github.com/volatiletech/sqlboiler/v4/types"
	"time"
)

type generateMessageIDRequest struct {
	UserID          int64  `json:"userId"`
	Channel         string `json:"channel"`
	MessageTemplate string `json:"messageTemplate"`
	Timestamp       int64  `json:"timestamp"`
	ExternalID      int64  `json:"externalId"`
}

type generateMessageIDResponse struct {
	ID string `json:"id"`
}

type sendMessageRequest struct {
	Params types.JSON `json:"params"`
}

type messageResponse struct {
	ID                string    `json:"id"`
	Status            string    `json:"status"`
	StatusDescription string    `json:"statusDescription"`
	StatusTime        time.Time `json:"statusTime"`
}

type operationStatus struct {
	Status            bool   `json:"status"`
	StatusDescription string `json:"statusDescription"`
}

type userChannelRequest struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	Channel   string `json:"channel"`
	Recipient string `json:"recipient"`
	CanNotify bool   `json:"canNotify"`
}

type userChannelUpdateRequest struct {
	Recipient string `json:"recipient"`
	CanNotify bool   `json:"canNotify"`
}

type userChannelResponse struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	Channel   string `json:"channel"`
	Recipient string `json:"recipient"`
	CanNotify bool   `json:"canNotify"`
}
