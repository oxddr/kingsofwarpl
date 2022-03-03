SELECT name,
       date,
       tabletop_url,
       count() AS players
  FROM Events e
       JOIN
       Results r ON e.id = r.event
 WHERE series = 'MP 2019'
 GROUP BY 1,
          2,
          3
 ORDER BY 2 DESC;



 
