db = kingsofwar-pl.sqlite3
data = data/events.json \
	data/liga2021_factions.json \
	data/liga2021_ranking.json \
	data/liga2021_results.json \
	data/liga2021_tournaments.json \
	data/liga2022_factions.json \
	data/liga2022_ranking.json \
	data/liga2022_results.json \
	data/liga2022_tournaments.json \
	data/players.json

db-restore: dump/kingsofwar-pl.sql
	rm -f $(db)
	sqlite3 $(db) < $<

db-dump:
	sqlite3 $(db) '.dump' > dump/kingsofwar-pl.sql

data-clean:
	rm -f $(data)

data-build: $(data)

$(data): data/%.json: sql/%.sql $(db)
	echo '.mode json' | cat - $< | sqlite3 $(db) | python -m json.tool > $@

gen-players: sql/players.sql
	rm -f content/player/*.md
	cat $< | sqlite3 $(db) | jq .[].tabletop_id | scripts/gen_players.sh

tttscraper: $(shell find tools/tttscraper2/ -name '*.go')
	go build -o $@ ./tools/tttscraper2
