-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS message_status
(
    id          bigserial PRIMARY KEY,
    message_id  varchar(255) NOT NULL REFERENCES message (id),
    status      varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    is_last     bool         NOT NULL DEFAULT false,
    created_at  timestamptz  NOT NULL DEFAULT now()
);

CREATE INDEX idx_message_status_status_created_at ON message_status (status, created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS message_status;
-- +goose StatementEnd
