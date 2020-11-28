UPDATE customers 
SET active = TRUE 
WHERE id = 2 RETURNING id, name, active, created;
