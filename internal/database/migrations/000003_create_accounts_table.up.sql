CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    balance FLOAT DEFAULT 0,
    currency_id INT REFERENCES currency(id) ON DELETE SET NULL
);
