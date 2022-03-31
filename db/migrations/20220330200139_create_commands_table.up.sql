-- Command reservations
CREATE TYPE command_response_type AS ENUM (
    'REQUESTED',
    'NOT_SUPPORTED', 
    'REJECTED', 
    'ACCEPTED', 
    'TIMEOUT', 
    'UNKNOWN_SESSION'
);

CREATE TABLE IF NOT EXISTS command_reservations (
    id             BIGSERIAL PRIMARY KEY,
    status         command_response_type NOT NULL,
    token_id       BIGINT NOT NULL,
    expiry_date    TIMESTAMP NOT NULL,
    reservation_id BIGSERIAL NOT NULL,
    location_id    TEXT NOT NULL,
    evse_uid       TEXT
);

ALTER TABLE command_reservations
    ADD CONSTRAINT fk_command_reservations_token_id
    FOREIGN KEY (token_id) 
    REFERENCES tokens(id) 
    ON DELETE RESTRICT;

-- Command starts
CREATE TABLE IF NOT EXISTS command_starts (
    id          BIGSERIAL PRIMARY KEY,
    status      command_response_type NOT NULL,
    token_id    BIGINT NOT NULL,
    location_id TEXT NOT NULL,
    evse_uid    TEXT
);

ALTER TABLE command_starts
    ADD CONSTRAINT fk_command_starts_token_id
    FOREIGN KEY (token_id) 
    REFERENCES tokens(id) 
    ON DELETE RESTRICT;

-- Command stops
CREATE TABLE IF NOT EXISTS command_stops (
    id         BIGSERIAL PRIMARY KEY,
    status     command_response_type NOT NULL,
    session_id TEXT NOT NULL
);

-- Command unlocks
CREATE TABLE IF NOT EXISTS command_unlocks (
    id           BIGSERIAL PRIMARY KEY,
    status       command_response_type NOT NULL,
    location_id  TEXT NOT NULL,
    evse_uid     TEXT NOT NULL,
    connector_id TEXT NOT NULL
);
