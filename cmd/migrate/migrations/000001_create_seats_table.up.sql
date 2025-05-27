CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    row CHAR(1) NOT NULL,
    col INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    UNIQUE (row, col)
)
