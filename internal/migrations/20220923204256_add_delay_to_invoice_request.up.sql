-- Invoice requests
ALTER TABLE invoice_requests
    ADD COLUMN release_date TIMESTAMP;
