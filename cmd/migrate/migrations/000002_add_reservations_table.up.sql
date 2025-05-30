CREATE TYPE reservation_status AS ENUM ('pending', 'confirmed', 'expired');

CREATE TABLE reservations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    status reservation_status NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE reserved_seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    reservation_id UUID NOT NULL REFERENCES reservations (id) ON DELETE CASCADE,
    seat_id UUID NOT NULL REFERENCES seats (id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    UNIQUE (reservation_id, seat_id)
);

CREATE INDEX idx_reserved_seats_reservation_id ON reserved_seats (reservation_id);

CREATE INDEX idx_reserved_seats_seat_id ON reserved_seats (seat_id);
