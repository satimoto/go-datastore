-- Routing events
CREATE TABLE IF NOT EXISTS routing_events (
    id                 BIGSERIAL PRIMARY KEY,
    currency           TEXT NOT NULL,
    currency_rate      BIGINT NOT NULL,
    currency_rate_msat BIGINT NOT NULL,
    incoming_chan_id   BIGINT NOT NULL,
    incoming_htlc_id   BIGINT NOT NULL,
    incoming_fiat      FLOAT NOT NULL,
    incoming_msat      BIGINT NOT NULL,
    outgoing_chan_id   BIGINT NOT NULL,
    outgoing_htlc_id   BIGINT NOT NULL,
    outgoing_fiat      FLOAT NOT NULL,
    outgoing_msat      BIGINT NOT NULL,
    fee_fiat           FLOAT NOT NULL,
    fee_msat           BIGINT NOT NULL,
    last_updated       TIMESTAMP NOT NULL
);
