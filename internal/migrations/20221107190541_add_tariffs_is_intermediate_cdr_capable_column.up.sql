-- Tariffs
ALTER TABLE tariffs
    ADD COLUMN is_intermediate_cdr_capable BOOLEAN NOT NULL DEFAULT true;
