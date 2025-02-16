CREATE TABLE income_records (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    account_id INT REFERENCES accounts(id) ON DELETE CASCADE,
    title VARCHAR NOT NULL,
    category_id INT REFERENCES income_categories(id) ON DELETE SET NULL,
    amount FLOAT NOT NULL,
    attachment VARCHAR
);
