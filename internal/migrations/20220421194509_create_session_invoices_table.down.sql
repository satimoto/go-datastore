-- Session invoices
ALTER TABLE IF EXISTS session_invoices 
    DROP CONSTRAINT IF EXISTS fk_session_invoices_session_id;

DROP TABLE IF EXISTS session_invoices;
