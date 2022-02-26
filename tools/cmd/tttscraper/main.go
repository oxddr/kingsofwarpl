package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oxddr/kingsofwarpl/tools/model"
	"go.uber.org/multierr"
)

const (
	PLAYER   = "Player"
	TP       = "TP"
	BONUS_TP = "Bonus TP"
	ATTR     = "AP/H2H"
	FACTION  = "Faction"
)

var (
	desc   = flag.String("description", "", "Path to file with league description")
	output = flag.String("output", "", "Path to file with league data.")

	idRegex = regexp.MustCompile("href=/profile/(.*)>Profile")

	columns = []string{PLAYER, TP, BONUS_TP, ATTR, FACTION}
)

func ExtractName(doc *goquery.Document) string {
	var name string
	doc.Find("#event-title").Each(func(i int, s *goquery.Selection) {
		name = strings.TrimSpace(s.Text())
	})
	return name
}

func ExtractDate(doc *goquery.Document) (time.Time, error) {
	var date string
	doc.Find("li[title='Event Date']").Each(func(i int, s *goquery.Selection) {
		date = strings.TrimSpace(s.Text())
	})
	return time.Parse("02-01-2006", date)
}

func ExtractLocation(doc *goquery.Document) string {
	var loc string
	doc.Find("li[title='Location']").Each(func(_ int, s *goquery.Selection) {
		loc = strings.Split(strings.TrimSpace(s.Text()), "/")[2]
	})
	return loc
}

func ExtractPlayers(doc *goquery.Document) ([]*model.Player, error) {
	indices := map[int]string{}
	doc.Find("#ladder thead tr").Each(func(i int, s *goquery.Selection) {
		s.Find("th").Each(func(i int, s *goquery.Selection) {
			for _, v := range columns {
				// log.Printf("Header %q, index %d", s.Text(), i)
				if v == s.Text() {
					indices[i] = v
					return
				}
			}
		})
	})

	var players []*model.Player
	var mErr error
	doc.Find("#ladder tbody tr").Each(func(_ int, s *goquery.Selection) {
		p := &model.Player{}
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			v, ok := indices[i]
			if !ok {
				return
			}

			var err error
			switch v {
			case PLAYER:
				a := s.Find("a")
				p.Name = strings.TrimSuffix(a.Text(), "R")
				hiddenA := a.AttrOr("title", "")

				matches := idRegex.FindStringSubmatch(hiddenA)
				if len(matches) > 1 {
					p.ID = matches[1]
				}
			case FACTION:
				if s.Text() != "-" {
					p.Faction = s.Text()
				}
			case TP:
				p.TP, err = strconv.Atoi(s.Text())
			case BONUS_TP:
				p.BonusTP, err = strconv.Atoi(s.Text())
			case ATTR:
				pts := strings.Split(s.Text(), "/")[0]
				p.AttritionPoints, err = strconv.Atoi(pts)
			}
			if err != nil {
				mErr = multierr.Append(mErr, err)
			}
		})
		players = append(players, p)
	})
	return players, mErr
}

func GetSingle(t *model.Tournament) (*model.TournamentResults, error) {
	res, err := http.Get(t.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to download %q: %v", t.URL, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status code error for %q: %d %s", t.URL, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize goquery: %v", err)
	}

	dt, err := ExtractDate(doc)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract tournament's date: %v", err)
	}

	players, err := ExtractPlayers(doc)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract tournament's players: %v", err)
	}
	rankPlayers(players)

	return &model.TournamentResults{
		Tournament: &model.Tournament{
			Name:     ExtractName(doc),
			Date:     dt,
			Location: ExtractLocation(doc),
			URL:      t.URL,
		},
		Players: players,
	}, nil
}

func rankPlayers(players []*model.Player) {
	sort.Slice(players, func(i, j int) bool {
		if players[i].TotalTP() != players[j].TotalTP() {
			return players[i].TotalTP() > players[j].TotalTP()
		}
		return players[i].AttritionPoints > players[j].AttritionPoints
	})

	for i, p := range players {
		p.Rank = i + 1
	}
}

func indexOf(s string, slice []string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	return -1
}

func main() {
	flag.Parse()

	if *desc == "" {
		log.Fatal("--league_description must be provided")
	}

	bytes, err := ioutil.ReadFile(*desc)
	if err != nil {
		log.Fatalf("Unable to read file with league description (%q): %v", *desc, err)
	}

	tournaments := []*model.Tournament{}
	if err := json.Unmarshal(bytes, &tournaments); err != nil {
		log.Fatalf("Unable to unmarshal league description: %v", err)
	}

	results := []*model.TournamentResults{}
	for _, t := range tournaments {
		r, err := GetSingle(t)
		if err != nil {
			log.Fatalf("Unable to get data for tournament %q: %v", t.Name, err)
		}
		results = append(results, r)
	}

	bytes, err = json.MarshalIndent(results, "", "\t")
	if err != nil {
		log.Fatalf("Unable to marshal league data: %v\n\n %v", err, results)
	}

	err = ioutil.WriteFile(*output, bytes, 0644)
	if err != nil {
		log.Fatalf("Unable to save data to file %q: %v", *output, err)
	}
}
