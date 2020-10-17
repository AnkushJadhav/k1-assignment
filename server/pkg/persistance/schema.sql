CREATE TABLE users (
    "id" CHAR(128) NOT NULL,
    "name" CHAR(100) NOT NULL,
    "email" CHAR(320) NOT NULL,
    "password" CHAR(20) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    "is_active" BOOLEAN DEFAULT true,
    "hits" INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);