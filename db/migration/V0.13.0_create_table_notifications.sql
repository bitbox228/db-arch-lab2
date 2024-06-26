CREATE TABLE IF NOT EXISTS notifications
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) NOT NULL,
    type    notification_type             NOT NULL,
    body_id INTEGER                       NOT NULL,
    time    TIMESTAMP                     NOT NULL
);