CREATE TABLE client_api
(
    id                INTEGER PRIMARY KEY,
    health_check_time INTEGER,
    toggle_client_api boolean ,
    listener_url      VARCHAR(200),
    url               VARCHAR(100),
    http_method       VARCHAR(100),
    header       VARCHAR(100),
    body              VARCHAR(100),
    interval_date;              dateTime
);
INSERT into client_api (id, health_check_time, toggle_client_api, listener_url, url, http_method, http_header, body,
                       interval_date)
VALUES (1, 3600, true, 'localhost:8080/listener-api', 'https://google.com', 'get', 'header', 'body', now());


# Queries on project
#select * from client_api where subdate(now(), interval health_check_time second) >= interval_date;