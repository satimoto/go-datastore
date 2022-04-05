-- name: SetElementRestrictionWeekday :exec
INSERT INTO element_restriction_weekdays 
    (element_restriction_id, weekday_id)
  VALUES ($1, $2);

-- name: UnsetElementRestrictionWeekdays :exec
DELETE FROM element_restriction_weekdays rw
  WHERE rw.element_restriction_id == $1;

-- name: ListElementRestrictionWeekdays :many
SELECT w.* FROM weekdays w
  INNER JOIN element_restriction_weekdays rw ON rw.weekday_id == w.id
  WHERE rw.element_restriction_id == $1
  ORDER BY w.id;
