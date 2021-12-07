package model

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type Player struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Rank            int    `json:"rank"`
	Faction         string `json:"faction"`
	TP              int    `json:"tp"`
	BonusTP         int    `json:"bonus_tp"`
	AttritionPoints int    `json:"attrition_points"`
}

type Tournament struct {
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
	URL      string    `json:"url"`
}

type TournamentResults struct {
	Tournament *Tournament `json:"tournament"`
	Players    []*Player   `json:"players"`
}

type League struct {
	Tournaments []*TournamentResults
}

func LeagueFromJSON(dir string) (*League, error) {
	l := &League{}

	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".json") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to list json files in %s: %v", dir, err)
	}

	for _, f := range files {
		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("unable to read %q: %v", f, err)
		}

		t := &TournamentResults{}
		err = json.Unmarshal(bytes, t)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal %q: %v", f, err)
		}
		log.Default().Printf("Building League from %s", f)
		l.Tournaments = append(l.Tournaments, t)
	}
	return l, nil
}

type RankedPlayer struct {
	Name        string    `json:"player"`
	ID          string    `json:"id"`
	Results     []*Result `json:"results"`
	Points      int       `json:"points"`
	PointsTotal int       `json:"points_total"`
	Rank        int       `json:"rank"`
}

type Result struct {
	Tournament *Tournament `json:"tournament"`
	Rank       int         `json:"rank"`
	Points     int         `json:"points"`
	Ranking    bool        `json:"is_ranking"`
}

func RankingFromJSON(file string) ([]*RankedPlayer, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q: %v", file, err)
	}

	ranking := []*RankedPlayer{}
	err = json.Unmarshal(bytes, &ranking)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal %q: %v", file, err)
	}
	return ranking, nil
}
