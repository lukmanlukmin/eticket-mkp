-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cards (
    id UUID PRIMARY KEY,
    card_number VARCHAR(50) NOT NULL UNIQUE,
    balance INTEGER NOT NULL DEFAULT 0,
    last_synced TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE terminals (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tariffs (
    id UUID PRIMARY KEY,
    entry_terminal_id UUID NOT NULL,
    exit_terminal_id UUID NOT NULL,
    fare INTEGER NOT NULL,
    effective_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_tariff_entry_terminal FOREIGN KEY (entry_terminal_id) REFERENCES terminals(id),
    CONSTRAINT fk_tariff_exit_terminal FOREIGN KEY (exit_terminal_id) REFERENCES terminals(id)
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    card_id UUID NOT NULL,
    entry_terminal_id UUID NOT NULL,
    exit_terminal_id UUID NOT NULL,
    fare INTEGER NOT NULL,
    entry_time TIMESTAMP NOT NULL,
    exit_time TIMESTAMP NOT NULL,
    source_device_id VARCHAR(100),
    synced_at TIMESTAMP,

    CONSTRAINT fk_trx_card FOREIGN KEY (card_id) REFERENCES cards(id),
    CONSTRAINT fk_trx_entry_terminal FOREIGN KEY (entry_terminal_id) REFERENCES terminals(id),
    CONSTRAINT fk_trx_exit_terminal FOREIGN KEY (exit_terminal_id) REFERENCES terminals(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS tariffs;
DROP TABLE IF EXISTS terminals;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
