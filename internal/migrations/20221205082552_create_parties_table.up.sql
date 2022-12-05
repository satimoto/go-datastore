
-- Parties
CREATE TABLE IF NOT EXISTS parties (
    id                          BIGSERIAL PRIMARY KEY,
    credential_id               BIGINT NOT NULL,
    country_code                TEXT NOT NULL,
    party_id                    TEXT NOT NULL,
    is_intermediate_cdr_capable BOOLEAN NOT NULL,
    publish_location            BOOLEAN NOT NULL,
    publish_null_tariff         BOOLEAN NOT NULL
);

ALTER TABLE parties 
    ADD CONSTRAINT uq_parties_country_code_party_id 
    UNIQUE (credential, country_code, party_id);

ALTER TABLE parties 
    ADD CONSTRAINT fk_parties_credential_id
    FOREIGN KEY (credential_id) 
    REFERENCES credentials(id) 
    ON DELETE CASCADE;

INSERT INTO parties (credential_id, country_code, party_id, is_intermediate_cdr_capable, publish_location, publish_null_tariff) VALUES 
    (34, "AT", "012", false, false, false), 
    (34, "AT", "120", false, false, false), 
    (34, "AT", "DAE", false, false, false), 
    (34, "AT", "ENA", false, false, false), 
    (34, "AT", "ETB", false, false, false), 
    (34, "AT", "HTB", false, false, false), 
    (34, "AT", "NEV", false, false, false), 
    (34, "AT", "POW", false, false, false), 
    (34, "AT", "STW", false, false, false), 
    (34, "AT", "SZG", false, false, false), 
    (34, "AT", "TLN", false, false, false), 
    (34, "AT", "TWG", false, false, false), 
    (34, "BE", "BLU", true, true, false), 
    (34, "BE", "OPT", false, false, false), 
    (34, "BE", "PWD", false, false, false), 
    (34, "BE", "TCB", false, false, false), 
    (34, "CH", "AMA", false, false, false), 
    (34, "CH", "EDH", false, false, false), 
    (34, "CH", "HPC", false, false, false), 
    (34, "CH", "SWI", true, true, false), 
    (34, "CZ", "LDL", false, false, false), 
    (34, "DE", "122", false, false, false), 
    (34, "DE", "444", false, false, false), 
    (34, "DE", "AUC", false, false, false), 
    (34, "DE", "AVI", false, false, false), 
    (34, "DE", "AVP", false, false, false), 
    (34, "DE", "BEM", false, false, false), 
    (34, "DE", "BOX", false, false, false), 
    (34, "DE", "BPE", false, false, false), 
    (34, "DE", "CCD", false, false, false), 
    (34, "DE", "CIW", false, false, false), 
    (34, "DE", "CON", false, false, false), 
    (34, "DE", "DES", false, false, false), 
    (34, "DE", "EDH", false, false, false), 
    (34, "DE", "EWE", false, false, false), 
    (34, "DE", "FKU", false, false, false), 
    (34, "DE", "HPC", false, false, false), 
    (34, "DE", "HUM", false, false, false), 
    (34, "DE", "JEG", false, false, false), 
    (34, "DE", "KDL", false, false, false), 
    (34, "DE", "LDL", false, false, false), 
    (34, "DE", "NAG", false, false, false), 
    (34, "DE", "ODR", false, false, false), 
    (34, "DE", "PFW", false, false, false), 
    (34, "DE", "POW", false, false, false), 
    (34, "DE", "SMC", false, false, false), 
    (34, "DE", "SRE", false, false, false), 
    (34, "DE", "SSW", false, false, false), 
    (34, "DE", "VHM", false, false, false), 
    (34, "DE", "WIT", false, false, false), 
    (34, "DK", "SPA", false, false, false), 
    (34, "EE", "ALX", false, false, false), 
    (34, "FR", "55C", false, false, false), 
    (34, "FR", "C01", true, true, false), 
    (34, "FR", "CG0", false, false, false), 
    (34, "FR", "CHA", true, true, false), 
    (34, "FR", "CPS", true, true, false), 
    (34, "FR", "DOU", false, false, false), 
    (34, "FR", "E04", true, true, false), 
    (34, "FR", "ECN", true, true, false), 
    (34, "FR", "EGR", true, true, false), 
    (34, "FR", "ELC", true, true, false), 
    (34, "FR", "GIR", true, true, false), 
    (34, "FR", "GLY", true, true, false), 
    (34, "FR", "IKA", true, true, false), 
    (34, "FR", "IPK", true, true, false), 
    (34, "FR", "IZM", true, true, false), 
    (34, "FR", "LDL", false, false, false), 
    (34, "FR", "LGC", true, true, false), 
    (34, "FR", "M06", true, true, false), 
    (34, "FR", "MAP", false, false, false), 
    (34, "FR", "OTH", true, true, false), 
    (34, "FR", "PA1", true, true, false), 
    (34, "FR", "S16", true, true, false), 
    (34, "FR", "S17", true, true, false), 
    (34, "FR", "S19", true, true, false), 
    (34, "FR", "S23", true, true, false), 
    (34, "FR", "S24", true, true, false), 
    (34, "FR", "S29", true, true, false), 
    (34, "FR", "S33", true, true, false), 
    (34, "FR", "S35", true, true, false), 
    (34, "FR", "S37", false, false, false), 
    (34, "FR", "S40", true, true, false), 
    (34, "FR", "S47", true, true, false), 
    (34, "FR", "S64", true, true, false), 
    (34, "FR", "S69", true, true, false), 
    (34, "FR", "S76", true, true, false), 
    (34, "FR", "S87", true, true, false), 
    (34, "FR", "SAE", true, true, false), 
    (34, "FR", "SDG", true, true, false), 
    (34, "FR", "SE1", true, true, false), 
    (34, "FR", "SIG", true, true, false), 
    (34, "FR", "SIP", false, false, false), 
    (34, "FR", "SLF", true, true, false), 
    (34, "FR", "SPS", true, true, false), 
    (34, "FR", "SSD", true, true, false), 
    (34, "FR", "SUA", true, true, false), 
    (34, "FR", "TCB", false, false, false), 
    (34, "FR", "TDA", true, true, false), 
    (34, "FR", "TNM", false, false, false), 
    (34, "FR", "URW", true, true, false), 
    (34, "FR", "V51", true, true, false), 
    (34, "FR", "V75", true, true, false), 
    (34, "GR", "EVZ", true, true, false), 
    (34, "HR", "PHR", false, false, false), 
    (34, "HR", "QLO", false, false, false), 
    (34, "HU", "AUD", false, false, false), 
    (34, "LU", "CHY", false, false, false), 
    (34, "LU", "SWO", false, false, false), 
    (34, "LU", "TCB", false, false, false), 
    (34, "LV", "CSD", false, false, false), 
    (34, "LV", "ELR", false, false, false), 
    (34, "LV", "ELX", true, true, false), 
    (34, "NL", "BPE", false, false, false), 
    (34, "NL", "CON", false, false, false), 
    (34, "NL", "CPI", true, true, false), 
    (34, "NL", "EVB", true, true, false), 
    (34, "NL", "TCB", false, false, false), 
    (34, "RO", "LDL", false, false, false), 
    (34, "RO", "LEK", false, false, false), 
    (34, "SE", "LDL", false, false, false), 
    (34, "SE", "PCS", false, false, false), 
    (34, "SK", "LDL", false, false, false);

-- Locations
ALTER TABLE locations
    ADD COLUMN is_intermediate_cdr_capable BOOLEAN NOT NULL DEFAULT true;

UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "012";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "120";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "DAE";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "ENA";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "ETB";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "HTB";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "NEV";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "POW";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "STW";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "SZG";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "TLN";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "AT" AND party_id = "TWG";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "BE" AND party_id = "BLU";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "BE" AND party_id = "OPT";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "BE" AND party_id = "PWD";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "BE" AND party_id = "TCB";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "CH" AND party_id = "AMA";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "CH" AND party_id = "EDH";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "CH" AND party_id = "HPC";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "CH" AND party_id = "SWI";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "CZ" AND party_id = "LDL";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "122";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "444";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "AUC";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "AVI";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "AVP";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "BEM";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "BOX";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "BPE";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "CCD";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "CIW";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "CON";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "DES";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "EDH";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "EWE";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "FKU";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "HPC";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "HUM";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "JEG";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "KDL";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "LDL";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "NAG";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "ODR";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "PFW";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "POW";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "SMC";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "SRE";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "SSW";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "VHM";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DE" AND party_id = "WIT";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "DK" AND party_id = "SPA";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "EE" AND party_id = "ALX";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "55C";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "C01";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "CG0";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "CHA";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "CPS";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "DOU";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "E04";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "ECN";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "EGR";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "ELC";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "GIR";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "GLY";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "IKA";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "IPK";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "IZM";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "LDL";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "LGC";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "M06";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "MAP";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "OTH";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "PA1";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S16";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S17";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S19";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S23";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S24";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S29";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S33";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S35";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "S37";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S40";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S47";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S64";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S69";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S76";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "S87";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SAE";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SDG";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SE1";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SIG";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "SIP";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SLF";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SPS";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SSD";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "SUA";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "TCB";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "TDA";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "FR" AND party_id = "TNM";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "URW";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "V51";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "FR" AND party_id = "V75";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "GR" AND party_id = "EVZ";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "HR" AND party_id = "PHR";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "HR" AND party_id = "QLO";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "HU" AND party_id = "AUD";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "LU" AND party_id = "CHY";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "LU" AND party_id = "SWO";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "LU" AND party_id = "TCB";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "LV" AND party_id = "CSD";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "LV" AND party_id = "ELR";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "LV" AND party_id = "ELX";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "NL" AND party_id = "BPE";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "NL" AND party_id = "CON";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "NL" AND party_id = "CPI";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "NL" AND party_id = "CPI";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "NL" AND party_id = "EVB";
UPDATE locations SET is_intermediate_cdr_capable = true WHERE country_code = "NL" AND party_id = "EVB";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "NL" AND party_id = "TCB";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "RO" AND party_id = "LDL";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "RO" AND party_id = "LEK";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "SE" AND party_id = "LDL";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "SE" AND party_id = "PCS";
UPDATE locations SET is_intermediate_cdr_capable = false WHERE country_code = "SK" AND party_id = "LDL";

-- Tariffs
ALTER TABLE tariffs
    DROP IF EXISTS is_intermediate_cdr_capable;
