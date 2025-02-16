CREATE TABLE installments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    account_id INT REFERENCES accounts(id) ON DELETE CASCADE,
    loan_id INT REFERENCES loans(id) ON DELETE CASCADE,
    title VARCHAR NOT NULL,
    amount FLOAT NOT NULL,
    date TIMESTAMP NOT NULL,
    scheduled_record_id INT REFERENCES scheduled_records(id) ON DELETE SET NULL
);
