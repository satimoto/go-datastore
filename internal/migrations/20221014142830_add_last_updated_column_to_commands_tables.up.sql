-- Command Reservations
ALTER TABLE command_reservations
    ADD COLUMN last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- Command Starts
ALTER TABLE command_starts
    ADD COLUMN last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- Command Stops
ALTER TABLE command_stops
    ADD COLUMN last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- Command Unlocks
ALTER TABLE command_unlocks
    ADD COLUMN last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
