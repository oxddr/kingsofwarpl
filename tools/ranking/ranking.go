package main

import (
	"encoding/json"
	"flag"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/oxddr/kingsofwarpl/tools"
)

var (
	resultsDir = flag.String("results_dir", "", "")
	output     = flag.String("output", "", "")
)

func ListAll(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".json") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func Rank2021(l *tools.League) []*tools.RankedPlayer {
	nameToPlayer := map[string]*tools.RankedPlayer{}

	for _, t := range l.Tournaments {
		for _, p := range t.Players {
			newPts := 30 - p.Rank + 1
			if newPts < 0 {
				newPts = 0
			}

			rankedPlayer, ok := nameToPlayer[p.Name]
			if !ok {
				rankedPlayer = &tools.RankedPlayer{
					Name: p.Name,
					ID:   p.ID,
				}
			}
			rankedPlayer.Results = append(rankedPlayer.Results, &tools.Result{
				Tournament: t.Tournament,
				Points:     newPts,
			})
			nameToPlayer[p.Name] = rankedPlayer
		}
	}

	topN := len(l.Tournaments) - 2
	if topN <= 0 {
		topN = 1
	}
	log.Printf("Counting top %d scores", topN)

	players := []*tools.RankedPlayer{}
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

	for i, r := range players {
		r.Rank = i + 1
	}

	return players
}

func main() {
	flag.Parse()

	files, err := ListAll(*resultsDir)
	if err != nil {
		log.Fatalf("Unable to list files from %q: %v", *resultsDir, err)
	}

	league, err := tools.LeagueFromJSON(files)
	if err != nil {
		log.Fatalf("Unable to build League from files: %v", err)
	}

	rankedPlayers := Rank2021(league)

	bytes, err := json.MarshalIndent(rankedPlayers, "", "\t")
	if err != nil {
		log.Fatalf("Unable to marshal ranking: %v", err)
	}
	err = ioutil.WriteFile(*output, bytes, 0644)
	if err != nil {
		log.Fatalf("Unable to write ranking to file %q: %v", *output, err)
	}

}
