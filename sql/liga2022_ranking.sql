WITH ranking_points AS (
    SELECT player,
           31 - rank() OVER (PARTITION BY event ORDER BY tp + bonus_tp DESC,
                attrition_points DESC) AS points,
           event, 
           1.0 * attrition_points / e.format as nap
      FROM Results r
           JOIN
           Events e ON r.event = e.id
     WHERE series = 'Liga 2022'
),
ranking_points_sorted AS (
    SELECT player,
           points,
           nap, 
           row_number() OVER (PARTITION BY player ORDER BY points DESC) AS result_rank
      FROM ranking_points
),
ranked_points AS (
    SELECT player,
           sum(points) AS ranked_points
      FROM ranking_points_sorted
     WHERE result_rank <= 2
     GROUP BY player
),
total_points AS (
    SELECT player,
           sum(points) AS total_points,
           sum(nap) as nap
      FROM ranking_points_sorted
     GROUP BY player
),
ranking_ranks AS (
    SELECT t.player AS id,
           name,
           row_number() OVER (ORDER BY ranked_points DESC,
           total_points DESC, nap desc) rank,
           ranked_points as points,
           total_points as tie_breaker1,
           printf("%.2f", nap) as tie_breaker2
      FROM ranked_points r
           JOIN
           total_points t ON r.player = t.player
           JOIN
           Players p ON r.player = p.tabletop_id
     ORDER BY rank ASC
)
SELECT *
  FROM ranking_ranks;

