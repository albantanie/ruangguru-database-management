-- TODO: answer here
UPDATE students
SET address = 'Bandung'
WHERE address IS NULL AND status = 'active';