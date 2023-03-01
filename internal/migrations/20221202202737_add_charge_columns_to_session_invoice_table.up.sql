-- Session invoices
ALTER TABLE session_invoices
    ADD COLUMN estimated_energy FLOAT NOT NULL DEFAULT 0;

ALTER TABLE session_invoices
    ADD COLUMN estimated_time FLOAT NOT NULL DEFAULT 0;

ALTER TABLE session_invoices
    ADD COLUMN metered_energy FLOAT NOT NULL DEFAULT 0;

ALTER TABLE session_invoices
    ADD COLUMN metered_time FLOAT NOT NULL DEFAULT 0;
