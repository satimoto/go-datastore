-- Psbt funding states
ALTER TABLE psbt_funding_states
    ADD COLUMN is_failed BOOLEAN NOT NULL DEFAULT false;
