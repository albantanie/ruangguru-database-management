-- TODO: answer here
SELECT 
    id, 
    CONCAT(first_name, ' ', last_name) AS student_name, 
    student_class, 
    final_score, 
    absent
FROM 
    reports
WHERE 
    final_score < 70 OR absent > 5
ORDER BY 
    id;