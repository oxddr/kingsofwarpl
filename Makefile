db = kingsofwar-pl.sqlite3
data = \
	data/events.json \
	data/events_2021.json \
	data/events_2022.json \
	data/factions_2021.json \
	data/factions_2022.json \
	data/players.json \
	data/ranking_2021.json \
	data/ranking_2022.json \
	data/results_2021.json \
	data/results_2022.json

db-restore: dump/kingsofwar-pl.sql
	rm -f $(db)
	sqlite3 $(db) < $<

db-dump: $(db)
	sqlite3 $(db) '.dump' > dump/kingsofwar-pl.sql

data-clean:
	rm -f data/*

data-build: $(data)

data/%.json: sql/%.sql $(db)
	echo '.mode json' | cat - $< | sqlite3 $(db) | python -m json.tool > $@

gen-players: sql/players.sql $(db)
	rm -f content/player/*.md
	cat $< | sqlite3 $(db) | jq .[].tabletop_id | scripts/gen_players.sh

tttscraper: $(shell find tools/tttscraper2/ -name '*.go')
	go build -o $@ ./tools/tttscraper2
