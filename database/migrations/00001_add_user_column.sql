-- +goose Up
-- +goose StatementBegin
CREATE TABLE user
(
    id          integer PRIMARY KEY AUTOINCREMENT, -- internal id
    provider_id text NOT NULL UNIQUE,              -- third-party id
    provider    text NOT NULL DEFAULT 'hanko',     -- For now, it is only Hanko
    name        text NOT NULL,
    email       text NOT NULL,
    status      text NOT NULL,                     -- Status of the user, can be a few values e.g. 'active', 'not-verified', 'banned', etc.
    created_at  datetime      DEFAULT CURRENT_TIMESTAMP,
    updated_at  datetime      DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
