CREATE DATABASE IF NOT EXISTS wallet;
USE wallet;

CREATE TABLE clients (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
    id CHAR(36) PRIMARY KEY,
    client_id CHAR(36) NOT NULL,
    balance DECIMAL(12,2) NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE transactions (
    id CHAR(36) PRIMARY KEY,
    account_id_from CHAR(36) NOT NULL,
    account_id_to CHAR(36) NOT NULL,
    amount DECIMAL(12,2) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id_from) REFERENCES accounts(id),
    FOREIGN KEY (account_id_to) REFERENCES accounts(id)
);

INSERT INTO clients (id, name, email) VALUES
  ('ba715ef5-3c2c-4ee9-ab49-ba564bddcafa', 'John Doe', 'john.doe@example.com'),
  ('4ab65695-e253-426c-916e-6295b1f4b009', 'Jane Doe', 'jane.doe@example.com');

INSERT INTO accounts (id, client_id, balance) VALUES
  ('546fbcb8-180a-4dd9-b36b-16304cf3e60a', 'ba715ef5-3c2c-4ee9-ab49-ba564bddcafa', 1000.00),
  ('ca88c60a-6092-49a7-9d58-6bd16fbc30d2', '4ab65695-e253-426c-916e-6295b1f4b009', 500.00);
