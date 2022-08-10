-- Tariff Price Components
ALTER TABLE IF EXISTS price_components 
    DROP CONSTRAINT IF EXISTS fk_tariff_price_components_step_rounding_id;

ALTER TABLE IF EXISTS price_components 
    DROP CONSTRAINT IF EXISTS fk_tariff_price_components_price_rounding_id;

ALTER TABLE IF EXISTS price_components 
    DROP CONSTRAINT IF EXISTS fk_tariff_price_components_element_id;

DROP TABLE IF EXISTS price_components;

DROP TYPE IF EXISTS tariff_dimension;

-- Price Component Rounding
DROP TABLE IF EXISTS price_component_roundings;

DROP TYPE IF EXISTS rounding_rule;

DROP TYPE IF EXISTS rounding_granularity;

-- Tariff Elements
ALTER TABLE IF EXISTS elements 
    DROP CONSTRAINT IF EXISTS fk_elements_element_restriction_id;

ALTER TABLE IF EXISTS elements 
    DROP CONSTRAINT IF EXISTS fk_elements_tariff_id;

DROP TABLE IF EXISTS elements;

-- Element Restriction Weekdays
ALTER TABLE IF EXISTS element_restriction_weekdays 
    DROP CONSTRAINT IF EXISTS fk_element_restriction_weekdays_weekday_id;

ALTER TABLE IF EXISTS element_restriction_weekdays 
    DROP CONSTRAINT IF EXISTS fk_element_restriction_weekdays_element_restriction_id;

DROP TABLE IF EXISTS element_restriction_weekdays;

-- Element Restrictions
DROP TABLE IF EXISTS element_restrictions;

-- Tariff Alt Texts
ALTER TABLE IF EXISTS tariff_alt_texts 
    DROP CONSTRAINT IF EXISTS fk_tariff_alt_texts_display_text_id;

ALTER TABLE IF EXISTS tariff_alt_texts 
    DROP CONSTRAINT IF EXISTS fk_tariff_alt_texts_tariff_id;

DROP TABLE IF EXISTS tariff_alt_texts;

-- Tariffs
ALTER TABLE IF EXISTS tariffs 
    DROP CONSTRAINT IF EXISTS fk_tariffs_energy_mix_id;
    
ALTER TABLE IF EXISTS tariffs 
    DROP CONSTRAINT IF EXISTS fk_tariffs_credential_id;

DROP INDEX IF EXISTS idx_tariffs_uid;

DROP TABLE IF EXISTS tariffs;

-- Tariff Restriction Weekdays
ALTER TABLE IF EXISTS tariff_restriction_weekdays 
    DROP CONSTRAINT IF EXISTS fk_tariff_restriction_weekdays_weekday_id;

ALTER TABLE IF EXISTS tariff_restriction_weekdays 
    DROP CONSTRAINT IF EXISTS fk_tariff_restriction_weekdays_tariff_restriction_id;

DROP TABLE IF EXISTS tariff_restriction_weekdays;

-- Tariff Restrictions
DROP TABLE IF EXISTS tariff_restrictions;

-- Day Of Week
DROP TABLE IF EXISTS weekdays;
