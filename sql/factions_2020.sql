WITH count_played AS (
    SELECT faction,
           count() AS played_count
      FROM Results r
           JOIN
           Events e ON r.event = e.id
     WHERE e.series = 'MP 2020'
     GROUP BY 1
),
per_player AS (
    SELECT player,
           faction
      FROM Results r
           JOIN
           Events e ON r.event = e.id
     WHERE e.series = 'MP 2020'
     GROUP BY 1,
              2
),
count_players AS (
    SELECT faction,
           count() AS player_count
      FROM per_player
     GROUP BY 1
)
SELECT p1.faction,
       played_count,
       player_count
  FROM count_played p1
       JOIN
       count_players p2 ON p1.faction = p2.faction
 ORDER BY 3 DESC,
          2 DESC;

;
