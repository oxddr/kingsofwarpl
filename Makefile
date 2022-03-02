tttscraper_bin = tools/cmd/tttscraper/tttscraper
scribe_bin = tools/cmd/scribe/scribe
players_bin = tools/cmd/players/players
db = kingsofwar-pl.sqlite3
data = data/events.json \
	data/liga2021_ranking.json \
	data/liga2021_tournaments.json \
	data/liga2021_factions.json \
	data/liga2022_ranking.json \
	data/liga2022_tournaments.json \
	data/liga2022_factions.json

db-restore: dump/kingsofwar-pl.sql
	rm -f $(db)
	sqlite3 $(db) < $<

db-dump:
	sqlite3 $(db) '.dump' > dump/kingsofwar-pl.sql

data-clean: $(data)
	rm $?

data-build: $(data)

$(data): data/%.json: sql/%.sql $(db)
	echo '.mode json' | cat - $< | sqlite3 $(db) | python -m json.tool > $@

.PHONY: player_pages
player_pages: data/liga2021/data.json data/liga2022/data.json $(players_bin)
	./$(players_bin) \
		--output=content/player \
		--leagues="data/liga2021/results.json" \
		--leagues="data/liga2022/results.json"

$(tttscraper_bin): tools/tttscraper/*.go tools/cmd/tttscraper/main.go
	cd tools/cmd/tttscraper && go build .

$(scribe_bin): tools/cmd/scribe/main.go tools/ranking/*.go tools/model/*.go
	cd tools/cmd/scribe/ && go build .

$(players_bin): tools/cmd/players/main.go
	cd tools/cmd/players && go build .
tttscraper: tools/cmd/tttscraper2/main.go
	go build -o $@ $<
