-- Session invoices
ALTER TABLE session_invoices
    DROP IF EXISTS metered_time;

ALTER TABLE session_invoices
    DROP IF EXISTS metered_energy;

ALTER TABLE session_invoices
    DROP IF EXISTS estimated_time;

ALTER TABLE session_invoices
    DROP IF EXISTS estimated_energy;
