SELECT  p.id,
        p.name,
        SUM(ss.total)
FROM candidate p   
JOIN (
    SELECT  sp.voter_id,
            v.candidate_id,
            CASE 
                WHEN COUNT(sp.voter_id) > 1 THEN 0 
                ELSE COUNT(sp.voter_id) 
            END total

    FROM votes sp
    LEFT JOIN votes v on sp.voter_id = v.voter_id 
    GROUP BY sp.voter_id,
             v.candidate_id
) ss ON p.id = ss.candidate_id
GROUP BY p.id,
        p.name
ORDER BY SUM(ss.total) DESC  
LIMIT 3

;