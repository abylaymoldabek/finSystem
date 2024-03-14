CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance NUMERIC NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    value NUMERIC NOT NULL,
    account_id INTEGER NOT NULL REFERENCES accounts(id),
    group_type VARCHAR(10) NOT NULL,
    account2_id INTEGER,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
