-- +goose Up
-- +goose StatementBegin
CREATE TABLE monitor
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT, -- internal id
    user_id    INTEGER NOT NULL,
    name       TEXT    NOT NULL,
    type       TEXT    NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE http_monitor
(
    monitor_id            INTEGER PRIMARY KEY NOT NULL,
    url                   TEXT                NOT NULL,
    interval_s            INTEGER             NOT NULL,
    retries               INTEGER             NOT NULL,
    timeout_s             INTEGER             NOT NULL,
    expected_status_codes TEXT                NOT NULL, -- values separated  by comma '200,300' or dash for a range, 200-299
    http_method           TEXT                NOT NULL,
    http_body             TEXT,
    http_headers          TEXT,
    FOREIGN KEY (monitor_id) REFERENCES monitor (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE monitor;
-- +goose StatementEnd
