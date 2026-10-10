-- +migrate Up
CREATE TABLE users
(
    id           int PRIMARY KEY AUTO_INCREMENT,
    name         varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL unique,
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE users;