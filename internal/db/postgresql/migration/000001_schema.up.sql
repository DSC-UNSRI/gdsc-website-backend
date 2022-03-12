CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    "uuid" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" VARCHAR NOT NULL,
    "birthdate" DATE NOT NULL
);