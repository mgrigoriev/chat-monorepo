CREATE TABLE IF NOT EXISTS participants (
    id SERIAL PRIMARY KEY,
    chatserver_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL
);
