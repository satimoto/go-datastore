-- Session updates
CREATE TABLE IF NOT EXISTS session_updates (
    id           BIGSERIAL PRIMARY KEY,
    session_id   BIGINT NOT NULL,
    user_id      BIGINT NOT NULL,
    kwh          FLOAT NOT NULL,
    total_cost   FLOAT,
    status       session_status_type NOT NULL,
    last_updated TIMESTAMP NOT NULL
);

ALTER TABLE session_updates 
    ADD CONSTRAINT fk_session_updates_session_id
    FOREIGN KEY (session_id) 
    REFERENCES sessions(id) 
    ON DELETE CASCADE;

ALTER TABLE session_updates 
    ADD CONSTRAINT fk_session_updates_user_id
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE;
