CREATE TABLE IF NOT EXISTS invites (
    id SERIAL PRIMARY KEY,
    chatserver_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL
);
