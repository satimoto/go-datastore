-- Routing events
ALTER TABLE IF EXISTS routing_events 
    DROP CONSTRAINT IF EXISTS fk_routing_events_node_id;

DROP TABLE IF EXISTS routing_events;

DROP TYPE IF EXISTS routing_event_type;
