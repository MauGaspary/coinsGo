CREATE TABLE accounts (
    account_id VARCHAR(255) PRIMARY KEY,
    balance    NUMERIC(20,4) NOT NULL DEFAULT 0.00
);

CREATE TABLE login_details (
    account_id VARCHAR(255) PRIMARY KEY,
    auth_token TEXT NOT NULL,
    CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES accounts(account_id)
);