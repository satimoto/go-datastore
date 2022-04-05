-- name: SetTariffRestrictionWeekday :exec
INSERT INTO tariff_restriction_weekdays 
    (tariff_restriction_id, weekday_id)
  VALUES ($1, $2);

-- name: UnsetTariffRestrictionWeekdays :exec
DELETE FROM tariff_restriction_weekdays rw
  WHERE rw.tariff_restriction_id = $1;

-- name: ListTariffRestrictionWeekdays :many
SELECT w.* FROM weekdays w
  INNER JOIN tariff_restriction_weekdays rw ON rw.weekday_id = w.id
  WHERE rw.tariff_restriction_id = $1
  ORDER BY w.id;
