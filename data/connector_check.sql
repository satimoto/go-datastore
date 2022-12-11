-- Connector check
SELECT l.country, l.uid, c.* FROM locations l
  LEFT JOIN evses e ON e.location_id = l.id
  LEFT JOIN connectors c ON c.evse_id = e.id
  WHERE country_code = 'NL' AND party_id = 'CPI';

-- Distinct countries for party
SELECT DISTINCT l.country FROM locations l
  WHERE country_code = 'NL' AND party_id = 'CPI'
  ORDER BY country ASC;

-- Count published by country
SELECT l.country, COUNT(l.country) AS cc FROM locations l
  WHERE l.is_published
  GROUP BY country
  ORDER BY cc DESC;

-- Added in the last week
SELECT l.country, l.uid, l.country_code, l.party_id, l.is_published, l.added_date, c.* FROM locations l
  LEFT JOIN evses e ON e.location_id = l.id
  LEFT JOIN connectors c ON c.evse_id = e.id
  WHERE l.added_date > NOW() - interval '1 week' AND c.tariff_id is not null AND l.is_published = false;

-- Set not removed unpublished locations to published
SELECT l.country, l.uid, e.status, c.* FROM locations l
  LEFT JOIN evses e ON e.location_id = l.id
  LEFT JOIN connectors c ON c.evse_id = e.id
  WHERE l.party_id not in ('V51', 'IPK', 'V75', 'TCB') AND e.status != 'REMOVED' AND 
    tariff_id is not null AND c.is_published and l.is_published = false;

UPDATE locations l SET is_published = true
  FROM evses e, connectors c
  WHERE e.location_id = l.id AND c.evse_id = e.id AND 
    l.party_id not in ('V51', 'IPK', 'V75', 'TCB') AND e.status != 'REMOVED' AND 
    tariff_id is not null AND c.is_published AND l.is_published = false;


-- Set not removed unpublished connectors to published
SELECT l.country, l.uid, e.status, c.* FROM locations l
  LEFT JOIN evses e ON e.location_id = l.id
  LEFT JOIN connectors c ON c.evse_id = e.id
  WHERE l.party_id not in ('V51', 'IPK', 'V75', 'TCB') AND e.status != 'REMOVED' AND 
    tariff_id is not null AND c.is_published = false;

UPDATE connectors c SET is_published = true, is_removed = false
  FROM evses e, locations l
  WHERE e.location_id = l.id AND c.evse_id = e.id AND 
    l.party_id not in ('V51', 'IPK', 'V75', 'TCB') AND e.status != 'REMOVED' AND 
    tariff_id is not null AND c.is_published = false;


-- Set connectors with missing tariffs to unpublished
SELECT c.* FROM connectors c
  LEFT JOIN tariffs t ON c.tariff_id = t.uid
  WHERE c.is_published = true AND (c.tariff_id is null OR t.uid is NULL)
  ORDER BY c.id;

UPDATE connectors c SET is_published = false 
  FROM connectors cj 
  LEFT JOIN tariffs tj ON cj.tariff_id = tj.uid
  WHERE c.id = cj.id AND cj.is_published = true AND (cj.tariff_id is null OR tj.uid is NULL)

-- Set location to unpublished where no connectors are published
SELECT l.id, l.uid, ec.e_count FROM locations l 
  LEFT JOIN (SELECT e.location_id, COUNT(e.location_id) AS e_count
			 FROM evses e 
			 LEFT JOIN connectors c ON e.id = c.evse_id WHERE c.is_published GROUP BY e.location_id) ec
    ON l.id = ec.location_id
	WHERE l.is_published AND ec.e_count is null


UPDATE locations l SET is_published = false 
  FROM locations lj 
  LEFT JOIN (SELECT e.location_id, COUNT(e.location_id) AS e_count
			 FROM evses e 
			 LEFT JOIN connectors c ON e.id = c.evse_id WHERE c.is_published GROUP BY e.location_id) ec
    ON lj.id = ec.location_id
	WHERE l.id = lj.id AND lj.is_published AND ec.e_count is null

