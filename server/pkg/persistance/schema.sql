CREATE TABLE users (
    id CHAR(128) NOT NULL,
    name CHAR(100) NOT NULL,
    email CHAR(255) NOT NULL,
    password CHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    hits INT NOT NULL,
    PRIMARY KEY (id)
);