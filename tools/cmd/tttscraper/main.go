package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/oxddr/kingsofwarpl/tools/model"
	"github.com/oxddr/kingsofwarpl/tools/tttscraper"
)

var (
	desc   = flag.String("description", "", "Path to file with league description")
	output = flag.String("output", "", "Path to file with league data.")
)

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
		r, err := tttscraper.Scrape(t.URL)
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
