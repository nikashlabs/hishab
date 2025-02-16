CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    photo_url VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    modified_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_active TIMESTAMP
);

CREATE UNIQUE INDEX users_email_unique ON users (email);

CREATE INDEX users_last_active_index ON users (last_active);