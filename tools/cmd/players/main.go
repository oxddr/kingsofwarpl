package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/oxddr/kingsofwarpl/tools/model"
	"gopkg.in/yaml.v3"
)

var (
	rankingPath = flag.String("ranking_path", "", "")
	output      = flag.String("output", "", "")

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

func main() {
	flag.Parse()

	ranking, err := model.RankingFromJSON(*rankingPath)
	if err != nil {
		log.Fatalf("Unable to create ranking from %q: %v", *rankingPath, err)
	}

	tpl, err := template.New("player").Parse(playerPageTmpl)
	if err != nil {
		log.Fatalf("Unable to create template: %v")
	}

	for _, p := range ranking {
		fm := &frontMatter{
			Title:  fmt.Sprintf("Gracz: %s", p.Name),
			Player: p.ID,
		}

		yamlBytes, err := yaml.Marshal(fm)
		if err != nil {
			log.Fatalf("Unable to create YAML: %v", err)
		}

		data := &tplData{
			FrontMatter: string(yamlBytes),
		}

		file := path.Join(*output, fmt.Sprintf("%s.md", p.ID))
		f, err := os.Create(file)
		if err != nil {
			log.Fatalf("Unable to create %q: %v", file, err)
		}
		defer f.Close()
		tpl.Execute(f, data)
		log.Printf("Created entry %s", file)
	}
}
