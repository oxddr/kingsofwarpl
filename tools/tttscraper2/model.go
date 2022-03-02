package main

import (
	"time"
)

type SingleResult struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
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
	Tournament *Tournament     `json:"tournament"`
	Players    []*SingleResult `json:"players"`
}
