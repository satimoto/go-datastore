-- Pending notifications
ALTER TABLE IF EXISTS pending_notifications 
    DROP CONSTRAINT IF EXISTS fk_pending_notifications_invoice_request_id;

ALTER TABLE IF EXISTS pending_notifications 
    DROP CONSTRAINT IF EXISTS fk_pending_notifications_node_id;

ALTER TABLE IF EXISTS pending_notifications 
    DROP CONSTRAINT IF EXISTS fk_pending_notifications_user_id;

DROP TABLE IF EXISTS pending_notifications;
