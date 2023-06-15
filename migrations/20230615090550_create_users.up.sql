CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id uuid default uuid_generate_v4 (),
    email varchar not null unique,
    username varchar not null unique,
    encrypted_password varchar not null
);