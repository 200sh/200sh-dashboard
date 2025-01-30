-- name: CreateHttpMonitor :one
INSERT INTO http_monitor(monitor_id,
                         url,
                         interval_s,
                         retries,
                         timeout_s,
                         expected_status_codes,
                         http_method,
                         http_body,
                         http_headers)
VALUES (?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?)
RETURNING *;

-- name: FindHttpMonitorByUserId :one
SELECT *
FROM http_monitor hm
JOIN main.monitor m on m.id = hm.monitor_id
WHERE m.user_id = ?;

-- name: UpdateHttpMonitor :one
UPDATE http_monitor
SET url = ?,
    interval_s = ?,
    retries = ?,
    timeout_s = ?,
    expected_status_codes = ?,
    http_method = ?,
    http_body = ?,
    http_headers = ?
WHERE monitor_id = ?
RETURNING *;