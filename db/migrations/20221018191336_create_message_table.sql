-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS message
(
    id          varchar(255) PRIMARY KEY,
    user_id     bigint      NOT NULL,
    external_id bigint      NOT NULL,
    channel     smallint    NOT NULL,
    template    smallint    NOT NULL,
    params      jsonb       NOT NULL DEFAULT '{}'::jsonb,
    timestamp   timestamptz NOT NULL
);

CREATE INDEX idx_message_user_id ON message (user_id);
CREATE INDEX idx_message_channel ON message (channel);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS message;
-- +goose StatementEnd
