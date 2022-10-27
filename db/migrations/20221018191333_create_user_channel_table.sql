-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_channel
(
    id         bigserial PRIMARY KEY,
    user_id    bigint       NOT NULL,
    channel    smallint     NOT NULL,
    recipient  varchar(255) NOT NULL,
    can_notify boolean      NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX uidx_user_channel_user_id_channel ON user_channel (user_id, channel);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_channel;
-- +goose StatementEnd
