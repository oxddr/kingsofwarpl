package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oxddr/kingsofwarpl/tools"
	"go.uber.org/multierr"
)

var (
	resultURL = flag.String("result_url", "", "Link to the event page on tabletop.to")
	output    = flag.String("output", "", "Path to output directory")

	idRegex = regexp.MustCompile("href=/profile/(.*)>Profile")
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

func ExtractPlayers(doc *goquery.Document) ([]*tools.Player, error) {
	var players []*tools.Player
	var mErr, err error
	doc.Find("#ladder tbody tr").Each(func(_ int, s *goquery.Selection) {
		p := &tools.Player{}
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				a := s.Find("a")
				p.Name = strings.TrimSuffix(a.Text(), "R")
				hiddenA := a.AttrOr("title", "")

				matches := idRegex.FindStringSubmatch(hiddenA)
				if len(matches) > 1 {
					p.ID = matches[1]
				}
			case 1:
				p.Rank, err = strconv.Atoi(s.Text())
				if err != nil {
					mErr = multierr.Append(mErr, err)
				}
			case 3:
				if s.Text() != "-" {
					p.Faction = s.Text()
				}
			case 4:
				p.TP, err = strconv.Atoi(s.Text())
				if err != nil {
					mErr = multierr.Append(mErr, err)
				}

			case 5:
				p.BonusTP, err = strconv.Atoi(s.Text())
				if err != nil {
					mErr = multierr.Append(mErr, err)
				}

			case 7:
				pts := strings.Split(s.Text(), "/")[0]
				p.AttritionPoints, err = strconv.Atoi(pts)
				if err != nil {
					mErr = multierr.Append(mErr, err)
				}

			}
		})
		players = append(players, p)
	})
	return players, mErr
}

func main() {
	flag.Parse()

	if *resultURL == "" {
		log.Fatal("--result_url must be provided")
	}

	res, err := http.Get(*resultURL)
	if err != nil {
		log.Fatalf("Unable to download %q: %v", *resultURL, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error for %q: %d %s", *resultURL, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Unable to initialize goquery: %v", err)
	}

	dt, err := ExtractDate(doc)
	if err != nil {
		log.Fatalf("Unable to extract tournament's date: %v", err)
	}

	players, err := ExtractPlayers(doc)
	if err != nil {
		log.Fatalf("Unable to extract tournament's players: %v", err)
	}

	t := &tools.TournamentResults{
		Tournament: &tools.Tournament{
			Name:     ExtractName(doc),
			Date:     dt,
			Location: ExtractLocation(doc),
			URL:      *resultURL,
		},
		Players: players,
	}

	bytes, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		log.Fatalf("Unable to marshal tournament data: %v\n\n %v", err, t)
	}

	err = ioutil.WriteFile(*output, bytes, 0644)
	if err != nil {
		log.Fatalf("Unable to save data to file %q: %v", *output, err)
	}

}
