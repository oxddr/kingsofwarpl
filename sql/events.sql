SELECT name,
       date,
       city,
       location,
       format,
       website AS event_url,
       tabletop_url AS signup_url
  FROM Events
 WHERE date > '2021-12-31'
 ORDER BY 2 DESC;
