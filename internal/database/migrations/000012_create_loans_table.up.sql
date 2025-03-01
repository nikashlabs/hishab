CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    account_id INT REFERENCES accounts(id) ON DELETE CASCADE,
    title VARCHAR NOT NULL,
    loan_type_id INT REFERENCES loan_types(id) ON DELETE SET NULL,
    amount FLOAT NOT NULL,
    date TIMESTAMP NOT NULL
);
