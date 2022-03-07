db = kingsofwar-pl.sqlite3

db-restore: dump/kingsofwar-pl.sql
	rm -f $(db)
	sqlite3 $(db) < $<

db-dump: $(db)
	sqlite3 $(db) '.dump' > dump/kingsofwar-pl.sql

data-clean:
	rm -f data/*

data-build: $(shell find sql -name "*.sql" | sed 's/sql/data/' | sed 's/sql/json/')

data/%.json: sql/%.sql $(db)
	echo '.mode json' | cat - $< | sqlite3 $(db) | python -m json.tool > $@

gen-players: sql/players.sql $(db)
	rm -f content/player/*.md
	echo '.mode json' | cat - $< | sqlite3 $(db) | jq -r .[].tabletop_id | scripts/gen_players.sh

tttscraper: $(shell find tools/tttscraper2/ -name '*.go')
	go build -o $@ ./tools/tttscraper2
