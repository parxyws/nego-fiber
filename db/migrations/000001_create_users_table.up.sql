CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY,
    username   VARCHAR(50) UNIQUE NOT NULL,
    email      VARCHAR(50) UNIQUE NOT NULL,
    password   VARCHAR(50) UNIQUE NOT NULL,

    full_name  VARCHAR(255)       NOT NULL,
    phone_num  VARCHAR(20)        NOT NULL,
    avatar_url TEXT,
    verified_at TIMESTAMP NUll,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS roles(
    id UUID PRIMARY KEY
)