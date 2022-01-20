package stats

import (
	"log"
	"sort"

	"bitbucket.org/creachadair/stringset"
	"github.com/oxddr/kingsofwarpl/tools/model"
)

type FactionStats struct {
	Name        string `json:"name"`
	TimesPlayed int    `json:"times_played"`
	NoPlayers   int    `json:"no_players"`
}

func CalculateFactions(league *model.League) []*FactionStats {
	timesPlayed := map[string]int{}
	noPlayers := map[string]stringset.Set{}

	for _, t := range league.Tournaments {
		for _, p := range t.Players {
			faction := p.Faction
			if faction == "" {
				faction = "(unknown)"
			}
			timesPlayed[faction]++

			set, ok := noPlayers[faction]
			if !ok {
				set = stringset.New()
				noPlayers[faction] = set
			}
			set.Add(p.Name)
		}
	}

	stats := []*FactionStats{}
	for f, timesPlayed := range timesPlayed {
		s := &FactionStats{
			Name:        f,
			NoPlayers:   len(noPlayers[f]),
			TimesPlayed: timesPlayed,
		}
		stats = append(stats, s)
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].TimesPlayed > stats[j].TimesPlayed
	})

	log.Printf("Recorded %d factions", len(stats))
	return stats
}
