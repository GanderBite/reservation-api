ALTER TABLE reservations
DROP CONSTRAINT fk_applied_discount,
DROP COLUMN applied_discount_code_id;

DROP TABLE discount_codes;
