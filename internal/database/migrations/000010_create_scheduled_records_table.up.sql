CREATE TABLE scheduled_records (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    account_id INT REFERENCES accounts(id) ON DELETE CASCADE,
    title VARCHAR NOT NULL,
    type VARCHAR CHECK (type IN ('expense', 'income', 'investment')),
    category_id INT REFERENCES expense_categories(id) ON DELETE SET NULL,
    amount FLOAT NOT NULL,
    status VARCHAR CHECK (status IN ('confirm', 'pending', 'cancel')),
    date TIMESTAMP NOT NULL,
    recurring_interval TIMESTAMP,
    recurring BOOLEAN DEFAULT FALSE
);
