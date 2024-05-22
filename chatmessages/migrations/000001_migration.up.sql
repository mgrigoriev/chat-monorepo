CREATE TABLE IF NOT EXISTS chatmessages (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    recipient_type SMALLINT NOT NULL,
    recipient_id BIGINT NOT NULL,
    content TEXT NOT NULL
);
