package http

import (
    "github.com/volatiletech/sqlboiler/v4/types"
    "time"
)

type generateMessageIDRequest struct {
    Channel         string `json:"channel"`
    UserID          int64  `json:"userId"`
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
