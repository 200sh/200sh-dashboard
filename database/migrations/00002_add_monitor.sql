-- +goose Up
-- +goose StatementBegin
CREATE TABLE monitor
(
    id         integer PRIMARY KEY AUTOINCREMENT, -- internal id
    user_id    integer NOT NULL,
    url        text    NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE monitor;
-- +goose StatementEnd
