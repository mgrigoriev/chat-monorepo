CREATE TABLE IF NOT EXISTS friendships (
     id SERIAL PRIMARY KEY,
     follower_id BIGINT NOT NULL,
     followed_id BIGINT NOT NULL,
     status VARCHAR(50) NOT NULL
);
