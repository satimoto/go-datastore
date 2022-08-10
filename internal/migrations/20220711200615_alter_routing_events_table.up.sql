-- Routing events
CREATE TYPE routing_event_status AS ENUM (
    'IN_FLIGHT', 
    'SETTLE',
    'FORWARD_FAIL',
    'LINK_FAIL'
);

ALTER TABLE routing_events
    ADD COLUMN event_status routing_event_status NOT NULL;

ALTER TABLE routing_events
    ADD COLUMN wire_failure INT;

ALTER TABLE routing_events
    ADD COLUMN failure_detail INT;

ALTER TABLE routing_events
    ADD COLUMN failure_string TEXT;
