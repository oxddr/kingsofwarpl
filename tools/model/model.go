package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func (p *Player) TotalTP() int {
	return p.TP + p.BonusTP
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

func LeagueFromJSON(file string) (*League, error) {
	l := &League{}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q: %v", file, err)
	}

	results := []*TournamentResults{}
	err = json.Unmarshal(bytes, &results)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal results from  %q: %v", file, err)
	}
	log.Default().Printf("Building League from %s", file)
	l.Tournaments = results

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
