-- Sessions
DELETE FROM pg_enum
  WHERE enumlabel = 'INVOICED'
    AND enumtypid = (
      SELECT oid FROM pg_type WHERE typname = 'session_status_type'
    );