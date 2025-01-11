-- +goose Up
-- +goose StatementBegin
CREATE TABLE user
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT, -- internal id
    provider_id TEXT    NOT NULL UNIQUE,           -- third-party id
    provider    TEXT    NOT NULL DEFAULT 'hanko',  -- For now, it is only Hanko
    name        TEXT    NOT NULL,
    email       TEXT    NOT NULL,
    status      INTEGER NOT NULL,                  -- Status of the user, can be a few values e.g. 'active', 'not-verified', 'banned', etc.
    created_at  DATETIME         DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME         DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
