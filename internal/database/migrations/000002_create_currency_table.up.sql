CREATE TABLE currency (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    conversion_rate FLOAT NOT NULL
);
