SELECT name,
       date,
       city,
       location,
       format,
       website AS event_url,
       tabletop_url AS signup_url
  FROM Events
 ORDER BY 2 DESC;
