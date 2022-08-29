-- Psbt funding states
ALTER TABLE IF EXISTS psbt_funding_states 
    DROP IF EXISTS is_failed;
