-- +migrate Up
CREATE TABLE users(
    Id text primary key,
    Name varchar(20) ,
    Email text UNIQUE,
    Password varchar(20),
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE products(
    Id text primary key,
    Name varchar(30),
    Price float,
    quantity int,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE carts(
    Id text primary key,
    Name varchar(30),
    Price float,
    quantity int,
    total float
);
-- +migrate Down
DROP TABLE users;
DROP TABLE products;
DROP TABLE cart;