CREATE TABLE harmony_config (
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(300) UNIQUE NOT NULL,
    config TEXT NOT NULL
);