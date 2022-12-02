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
  WHERE l.publish
  GROUP BY country
  ORDER BY cc DESC;