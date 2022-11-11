-- Has to be tariffs
CREATE TABLE IF NOT EXISTS htb_tariffs (
    id                BIGSERIAL PRIMARY KEY,
    name              TEXT NOT NULL,
    currency          TEXT NOT NULL,
    time_price        FLOAT,
    time_min_duration INTEGER,
    energy_price      FLOAT,
    flat_price        FLOAT,
    last_updated      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE htb_tariffs ADD CONSTRAINT uq_htb_tariffs_name UNIQUE (name);
