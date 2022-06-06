CREATE TABLE ClientApi
(
    id                INTEGER PRIMARY KEY,
    health_check_time INTEGER,
    toggle_client_api boolean ,
    listener_url      VARCHAR(200),
    url               VARCHAR(100),
    http_method       VARCHAR(100),
    http_headey       VARCHAR(100),
    body              VARCHAR(100),
    date_interval              dateTime
);
INSERT into ClientApi (id, health_check_time, toggle_client_api, listener_url, url, http_method, http_headey, body,
                       date_interval)
VALUES (1, 3600, true, 'localhost:8080/listener-api', 'https://google.com', 'get', 'header', 'body', now());

# select * from ClientApi where subdate(now(), interval health_check_time second) >= date_interval;