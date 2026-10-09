CREATE TABLE users
(
    id           int PRIMARY KEY AUTO_INCREMENT,
    name         varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL unique,
    password     text NOT NULL,
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE users;