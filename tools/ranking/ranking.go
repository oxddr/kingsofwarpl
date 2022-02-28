package ranking

import (
	"fmt"
	"log"
	"sort"

	"github.com/oxddr/kingsofwarpl/tools/model"
)

type tournamentSelector int

const (
	N_MINUS_2 tournamentSelector = iota
	ALL
)

func Build(year int, league *model.League) ([]*model.RankedPlayer, error) {
	switch year {
	case 2021:
		return rank2021(league), nil
	case 2022:
		return rank2022(league), nil
	}
	return nil, fmt.Errorf("No rankign algorithm specified for %d", year)

}

func rank2022(l *model.League) []*model.RankedPlayer {
	return rank(l, ALL)
}

func rank2021(l *model.League) []*model.RankedPlayer {
	return rank(l, N_MINUS_2)
}

func (t tournamentSelector) Top(n int) int {
	switch t {
	case N_MINUS_2:
		top := n - 2
		if top <= 0 {
			return 1
		}
		return top
	}
	return n
}

func rank(l *model.League, t tournamentSelector) []*model.RankedPlayer {
	nameToPlayer := map[string]*model.RankedPlayer{}

	for _, t := range l.Tournaments {
		for _, p := range t.Players {
			basePoints := 30
			if len(t.Players) > 30 {
				basePoints = len(t.Players)
			}

			newPts := basePoints - p.Rank + 1
			if newPts < 0 {
				newPts = 0
			}

			rankedPlayer, ok := nameToPlayer[p.Name]
			if !ok {
				rankedPlayer = &model.RankedPlayer{
					Name: p.Name,
					ID:   p.ID,
				}
			}
			rankedPlayer.Results = append(rankedPlayer.Results, &model.Result{
				Tournament: t.Tournament,
				Points:     newPts,
				Rank:       p.Rank,
			})
			nameToPlayer[p.Name] = rankedPlayer
		}
	}

	topN := t.Top(len(l.Tournaments))
	log.Printf("Counting top %d scores", topN)

	players := []*model.RankedPlayer{}
	for _, player := range nameToPlayer {
		sort.Slice(player.Results, func(i, j int) bool {
			return player.Results[i].Points > player.Results[j].Points
		})
		for i, result := range player.Results {
			if i < topN {
				result.Ranking = true
				player.Points += result.Points
			}
			player.PointsTotal += result.Points
		}
		players = append(players, player)

		sort.Slice(player.Results, func(i, j int) bool {
			return player.Results[i].Tournament.Date.Before(player.Results[j].Tournament.Date)
		})

	}

	sort.Slice(players, func(i int, j int) bool {
		if players[i].Points == players[j].Points {
			return players[i].PointsTotal > players[j].PointsTotal
		}
		return players[i].Points > players[j].Points
	})

	players[0].Rank = 1
	for i := 1; i < len(players); i++ {
		if players[i].Points == players[i-1].Points && players[i].PointsTotal == players[i-1].PointsTotal {
			players[i].Rank = players[i-1].Rank
		} else {
			players[i].Rank = players[i-1].Rank + 1
		}
	}

	return players
}