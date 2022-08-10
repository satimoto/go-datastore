-- Command unlocks
DROP TABLE IF EXISTS command_unlocks;

-- Command stops
DROP TABLE IF EXISTS command_stops;

-- Command starts
ALTER TABLE IF EXISTS command_starts 
    DROP CONSTRAINT IF EXISTS fk_command_starts_token_id;

DROP TABLE IF EXISTS command_starts;

-- Command reservations
ALTER TABLE IF EXISTS command_reservations 
    DROP CONSTRAINT IF EXISTS fk_command_reservations_token_id;

DROP TABLE IF EXISTS command_reservations;

DROP TYPE IF EXISTS command_response_type;
