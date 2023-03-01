-- Pending notifications
CREATE TABLE IF NOT EXISTS pending_notifications (
    id                 BIGSERIAL PRIMARY KEY,
    user_id            BIGINT NOT NULL,
    node_id            BIGINT NOT NULL,
    invoice_request_id BIGINT,
    device_token       TEXT NOT NULL,
    type               TEXT NOT NULL,
    send_date          TIMESTAMPTZ NOT NULL
);

ALTER TABLE pending_notifications 
    ADD CONSTRAINT fk_pending_notifications_user_id
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE;

ALTER TABLE pending_notifications 
    ADD CONSTRAINT fk_pending_notifications_node_id
    FOREIGN KEY (node_id) 
    REFERENCES nodes(id) 
    ON DELETE CASCADE;

ALTER TABLE pending_notifications 
    ADD CONSTRAINT fk_pending_notifications_invoice_request_id
    FOREIGN KEY (invoice_request_id) 
    REFERENCES invoice_requests(id) 
    ON DELETE CASCADE;
