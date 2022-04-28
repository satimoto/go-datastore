-- Session invoices
CREATE TABLE IF NOT EXISTS session_invoices (
    id              BIGSERIAL PRIMARY KEY,
    session_id      BIGINT NOT NULL,
    amount_fiat     FLOAT NOT NULL,
    amount_msat     BIGINT NOT NULL,
    commission_fiat FLOAT NOT NULL,
    commission_msat BIGINT NOT NULL,
    tax_fiat        FLOAT NOT NULL,
    tax_msat        BIGINT NOT NULL,
    currency        TEXT NOT NULL,
    payment_request TEXT NOT NULL,
    settled         BOOLEAN NOT NULL,
    expired         BOOLEAN NOT NULL,
    last_updated    TIMESTAMP NOT NULL
);

ALTER TABLE session_invoices
    ADD CONSTRAINT fk_session_invoices_session_id
    FOREIGN KEY (session_id) 
    REFERENCES sessions(id) 
    ON DELETE RESTRICT;
