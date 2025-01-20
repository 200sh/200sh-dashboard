-- +goose Up
-- +goose StatementBegin
CREATE TABLE monitor
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT, -- internal id
    user_id    INTEGER NOT NULL,
    url        TEXT    NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE monitor;
-- +goose StatementEnd
