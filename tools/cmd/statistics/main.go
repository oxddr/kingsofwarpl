package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"bitbucket.org/creachadair/stringset"
	"github.com/oxddr/kingsofwarpl/tools/model"
)

const (
	tournaments = "tournaments"
	factions    = "factions"
)

var (
	resultsDir = flag.String("results_dir", "", "")
	output     = flag.String("output", "", "")
	mode       = flag.String("mode", tournaments, "")
)

type FactionStat struct {
	Name        string `json:"name"`
	TimesPlayed int    `json:"times_played"`
	NoPlayers   int    `json:"no_players"`
}

type TournamentStat struct {
	model.Tournament
	NoPlayers int `json:"no_players"`
}

func main() {
	flag.Parse()

	league, err := model.LeagueFromJSON(*resultsDir)
	if err != nil {
		log.Fatalf("Unable to build League model frm %s: %v", *resultsDir, err)
	}

	var v interface{}
	if *mode == tournaments {
		tStat := []*TournamentStat{}
		for _, t := range league.Tournaments {
			stat := &TournamentStat{
				Tournament: *t.Tournament,
				NoPlayers:  len(t.Players),
			}
			tStat = append(tStat, stat)
		}
		sort.Slice(tStat, func(i, j int) bool {
			return tStat[i].Date.Before(tStat[j].Date)
		})
		v = tStat
	} else if *mode == factions {
		v = calculateFactions(league)
	} else {
		log.Fatalf("Unknown mode %q", *mode)
	}

	if err := saveJSON(v, *output); err != nil {
		log.Fatalf("Unable to save statistics as JSON to file: %v", err)
	}
}

func calculateFactions(league *model.League) []*FactionStat {
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

	stats := []*FactionStat{}
	for f, timesPlayed := range timesPlayed {
		s := &FactionStat{
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

func saveJSON(v interface{}, file string) error {
	bytes, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Errorf("unable to marshal data: %v", err)
	}
	err = ioutil.WriteFile(file, bytes, 0644)
	if err != nil {
		return fmt.Errorf("unable to write data to file %q: %v", *output, err)
	}
	return nil
}
