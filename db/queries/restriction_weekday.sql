-- name: SetRestrictionWeekday :exec
INSERT INTO restriction_weekdays (restriction_id, weekday_id)
  VALUES ($1, $2);

-- name: UnsetRestrictionWeekdays :exec
DELETE FROM restriction_weekdays rw
  WHERE rw.restriction_id == $1;

-- name: ListRestrictionWeekdays :many
SELECT w.* FROM weekdays w
  INNER JOIN restriction_weekdays rw ON rw.weekday_id == w.id
  WHERE rw.restriction_id == $1
  ORDER BY w.id;
