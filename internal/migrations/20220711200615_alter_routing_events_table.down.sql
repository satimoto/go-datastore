-- Routing events
ALTER TABLE IF EXISTS routing_events 
    DROP IF EXISTS failure_string;

ALTER TABLE IF EXISTS routing_events 
    DROP IF EXISTS failure_detail;

ALTER TABLE IF EXISTS routing_events 
    DROP IF EXISTS wire_failure;

ALTER TABLE IF EXISTS routing_events 
    DROP IF EXISTS event_status;

DROP TYPE IF EXISTS routing_event_status;
