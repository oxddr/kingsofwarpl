package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/spf13/pflag"

	"github.com/oxddr/kingsofwarpl/tools/model"
	"gopkg.in/yaml.v3"
)

var (
	leagues = pflag.StringSlice("leagues", []string{}, "")
	output  = pflag.String("output", "", "")

	playerPageTmpl = `---
{{ .FrontMatter }}---
`
)

type frontMatter struct {
	Title  string `yaml:"title"`
	Player string `yaml:"player"`
}

type tplData struct {
	FrontMatter string
}

func ExtractPlayers(leaguePaths []string) (map[string]*frontMatter, error) {
	players := map[string]*frontMatter{}
	for _, leaguePath := range leaguePaths {
		league, err := model.LeagueFromJSON(leaguePath)
		if err != nil {
			return nil, fmt.Errorf("unable to create league from %q: %v", *&leaguePath, err)
		}

		for _, t := range league.Tournaments {
			for _, p := range t.Players {
				players[p.ID] = &frontMatter{
					Title:  fmt.Sprintf("Gracz: %s", p.Name),
					Player: p.ID,
				}
			}
		}
	}
	return players, nil
}

func main() {
	pflag.Parse()

	tpl, err := template.New("player").Parse(playerPageTmpl)
	if err != nil {
		log.Fatalf("Unable to create template: %v")
	}

	playersFM, err := ExtractPlayers(*leagues)
	if err != nil {
		log.Fatalf("Unable to extract players: %v", err)
	}

	for _, fm := range playersFM {
		yamlBytes, err := yaml.Marshal(fm)
		if err != nil {
			log.Fatalf("Unable to create YAML: %v", err)
		}

		data := &tplData{
			FrontMatter: string(yamlBytes),
		}

		file := path.Join(*output, fmt.Sprintf("%s.md", fm.Player))
		f, err := os.Create(file)
		if err != nil {
			log.Fatalf("Unable to create %q: %v", file, err)
		}
		defer f.Close()
		tpl.Execute(f, data)
		log.Printf("Created entry %s", file)
	}
}
