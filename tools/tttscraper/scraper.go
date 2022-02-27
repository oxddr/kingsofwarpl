package tttscraper

import (
	"fmt"
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
	idRegex = regexp.MustCompile("href=/profile/(.*)>Profile")
	columns = []string{PLAYER, TP, BONUS_TP, ATTR, FACTION}
)

func Scrape(url string) (*model.TournamentResults, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to download %q: %v", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error for %q: %d %s", url, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize goquery: %v", err)
	}

	dt, err := extractDate(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to extract tournament's date: %v", err)
	}

	players, err := extractPlayers(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to extract tournament's players: %v", err)
	}
	rankPlayers(players)

	return &model.TournamentResults{
		Tournament: &model.Tournament{
			Name:     extractName(doc),
			Date:     dt,
			Location: extractLocation(doc),
			URL:      url,
		},
		Players: players,
	}, nil
}

func extractName(doc *goquery.Document) string {
	var name string
	doc.Find("#event-title").Each(func(i int, s *goquery.Selection) {
		name = strings.TrimSpace(s.Text())
	})
	return name
}

func extractDate(doc *goquery.Document) (time.Time, error) {
	var date string
	doc.Find("li[title='Event Date']").Each(func(i int, s *goquery.Selection) {
		date = strings.TrimSpace(s.Text())
	})
	return time.Parse("02-01-2006", date)
}

func extractLocation(doc *goquery.Document) string {
	var loc string
	doc.Find("li[title='Location']").Each(func(_ int, s *goquery.Selection) {
		loc = strings.Split(strings.TrimSpace(s.Text()), "/")[2]
	})
	return loc
}

func extractPlayers(doc *goquery.Document) ([]*model.Player, error) {
	indices := map[int]string{}
	doc.Find("#ladder thead tr").Each(func(i int, s *goquery.Selection) {
		s.Find("th").Each(func(i int, s *goquery.Selection) {
			for _, v := range columns {
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
