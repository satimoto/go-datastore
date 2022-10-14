-- Command Unlocks
ALTER TABLE command_unlocks
    DROP IF EXISTS last_updated;

-- Command Stops
ALTER TABLE command_stops
    DROP IF EXISTS last_updated;

-- Command Starts
ALTER TABLE command_starts
    DROP IF EXISTS last_updated;

-- Command Reservations
ALTER TABLE command_reservations
    DROP IF EXISTS last_updated;
