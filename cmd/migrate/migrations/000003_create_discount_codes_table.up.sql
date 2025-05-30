CREATE TABLE IF NOT EXISTS discount_codes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    code TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    price NUMERIC(10, 2) NOT NULL
);

ALTER TABLE reservations
ADD COLUMN applied_discount_code_id UUID,
ADD CONSTRAINT fk_applied_discount FOREIGN KEY (applied_discount_code_id) REFERENCES discount_codes (id) ON DELETE SET NULL;
