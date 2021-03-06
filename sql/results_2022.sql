WITH all_results AS (
    SELECT tabletop_id AS id,
           p.name,
           e.tabletop_url,
           e.name AS event_name,
           e.date AS event_date,
           rank() OVER (PARTITION BY e.id ORDER BY tp + bonus_tp DESC,
           attrition_points DESC) AS event_rank,
           best_of
      FROM Results r
           JOIN
           Players p ON r.player = p.tabletop_id
           JOIN
           Events e ON e.id = r.event
           JOIN
           Series s ON e.series = s.name
     WHERE series = 'Liga 2022'
),
results_with_ranks AS (
    SELECT *,
           31 - event_rank AS points,
           rank() OVER (PARTITION BY id ORDER BY event_rank ASC) AS rank
      FROM all_results
)
SELECT id,
       name,
       tabletop_url,
       event_name,
       event_date,
       event_rank,
       points,
       rank,
       IIF(rank <= best_of, 1, 0) AS is_ranked
  FROM results_with_ranks
 ORDER BY event_date DESC;
