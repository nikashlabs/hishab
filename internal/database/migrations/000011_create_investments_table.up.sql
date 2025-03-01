CREATE TABLE investments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    account_id INT REFERENCES accounts(id) ON DELETE CASCADE,
    title VARCHAR NOT NULL,
    investment_type_id INT REFERENCES investment_types(id) ON DELETE SET NULL,
    amount FLOAT NOT NULL,
    date TIMESTAMP NOT NULL,
    scheduled_record_id INT REFERENCES scheduled_records(id) ON DELETE SET NULL
);
