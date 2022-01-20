package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/oxddr/kingsofwarpl/tools/model"
	"github.com/oxddr/kingsofwarpl/tools/ranking"
	"github.com/oxddr/kingsofwarpl/tools/stats"
)

type LeagueStats struct {
	Ranking     []*model.RankedPlayer      `json:"ranking"`
	Factions    []*stats.FactionStats      `json:"factions"`
	Tournaments []*model.TournamentResults `json:"tournaments"`
}

var (
	resultsFile = flag.String("results", "", "Path to file with results")
	year        = flag.Int("year", 0, "Year for the statistics, can influence selected algorithms.")
	output      = flag.String("output", "", "Path to file with generated statistics")
)

func main() {
	flag.Parse()

	league, err := model.LeagueFromJSON(*resultsFile)
	if err != nil {
		log.Fatalf("Unable to build League model from %s: %v", *resultsFile, err)
	}

	r, err := ranking.Build(*year, league)
	if err != nil {
		log.Fatalf("Unable to build ranking for year %d: %v", *year, err)
	}

	stats := &LeagueStats{
		Tournaments: league.Tournaments,
		Factions:    stats.CalculateFactions(league),
		Ranking:     r,
	}

	bytes, err := json.MarshalIndent(stats, "", "\t")
	if err != nil {
		log.Fatalf("unable to marshal data: %v", err)
	}
	err = ioutil.WriteFile(*output, bytes, 0644)
	if err != nil {
		log.Fatalf("unable to write data to file %q: %v", *output, err)
	}
}
