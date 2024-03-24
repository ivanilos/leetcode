# Write your MySQL query statement below

SELECT x, y, z, 
    CASE WHEN
        GREATEST(x, y, z) >= LEAST(x, y, z) + LEAST(GREATEST(x, y), GREATEST(y, z), GREATEST(x, z)) 
        THEN "No"
    ELSE "Yes" END
    AS "triangle"
FROM Triangle
