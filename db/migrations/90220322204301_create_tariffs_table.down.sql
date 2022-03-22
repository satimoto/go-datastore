-- Tariff Price Components
ALTER TABLE IF EXISTS price_components 
    DROP CONSTRAINT IF EXISTS fk_tariff_price_components_element_id;

DROP TABLE IF EXISTS price_components;

DROP TYPE IF EXISTS location_type;

-- Tariff Elements
ALTER TABLE IF EXISTS elements 
    DROP CONSTRAINT IF EXISTS fk_elements_restriction_id;

ALTER TABLE IF EXISTS elements 
    DROP CONSTRAINT IF EXISTS fk_elements_tariff_id;

DROP TABLE IF EXISTS elements;

-- Tariff Restriction Weekdays
ALTER TABLE IF EXISTS restriction_weekdays 
    DROP CONSTRAINT IF EXISTS fk_restriction_weekdays_weekday_id;

ALTER TABLE IF EXISTS restriction_weekdays 
    DROP CONSTRAINT IF EXISTS fk_restriction_weekdays_restriction_id;

DROP TABLE IF EXISTS restriction_weekdays;

-- Day Of Week
DROP TABLE IF EXISTS weekdays;

-- Tariff Restrictions
DROP TABLE IF EXISTS restrictions;

-- Tariff Alt Texts
ALTER TABLE IF EXISTS tariff_alt_texts 
    DROP CONSTRAINT IF EXISTS fk_tariff_alt_texts_display_text_id;

ALTER TABLE IF EXISTS tariff_alt_texts 
    DROP CONSTRAINT IF EXISTS fk_tariff_alt_texts_tariff_id;

DROP TABLE IF EXISTS tariff_alt_texts;

-- Tariffs
ALTER TABLE IF EXISTS tariffs 
    DROP CONSTRAINT IF EXISTS fk_tariffs_energy_mix_id;

DROP INDEX IF EXISTS idx_tariffs_uid;

DROP TABLE IF EXISTS tariffs;
