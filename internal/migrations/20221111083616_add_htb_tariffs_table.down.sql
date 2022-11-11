-- Has to be tariffs
ALTER TABLE IF EXISTS htb_tariffs 
    DROP CONSTRAINT IF EXISTS uq_htb_tariffs_name;

DROP TABLE IF EXISTS htb_tariffs;
