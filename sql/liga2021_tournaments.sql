SELECT name,
       date,
       website,
       count() AS players
  FROM Events e
       JOIN
       Results r ON e.id = r.event
 WHERE series = 'Liga 2021'
 GROUP BY 1,
          2,
          3
 ORDER BY 2 DESC;



 
