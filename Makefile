tttscraper_bin = tools/cmd/tttscraper/tttscraper
scribe_bin = tools/cmd/scribe/scribe
players_bin = tools/cmd/players/players

data/events.json:
	curl -L "https://docs.google.com/spreadsheets/d/e/2PACX-1vQDorbeJ3svVWrbYYvtCNBOLif-mCCtjz45ndjxUtF0Ec_o77D20E5ejQtPcvM-YguvU1wH6BTxCaoC/pub?gid=790449776&single=true&output=csv" |  mlr --csv sort -f date  | mlr --c2j --jlistwrap cat > $@


data/liga2021/description.json:
	curl -L "https://docs.google.com/spreadsheets/d/e/2PACX-1vQDorbeJ3svVWrbYYvtCNBOLif-mCCtjz45ndjxUtF0Ec_o77D20E5ejQtPcvM-YguvU1wH6BTxCaoC/pub?gid=0&single=true&output=csv" |  mlr --csv sort -f date  | mlr --c2j --jlistwrap cat > $@

data/liga2021/results.json: data/liga2021/description.json $(tttscraper_bin)
	./$(tttscraper_bin) --description $<  --output $@

data/liga2021/data.json: data/liga2021/results.json $(scribe_bin)
	./$(scribe_bin) --year 2021 --results $< --output $@

data/liga2022/description.json:
	curl -L "https://docs.google.com/spreadsheets/d/e/2PACX-1vQDorbeJ3svVWrbYYvtCNBOLif-mCCtjz45ndjxUtF0Ec_o77D20E5ejQtPcvM-YguvU1wH6BTxCaoC/pub?gid=1575906467&single=true&output=csv" |  mlr --csv sort -f date  | mlr --c2j --jlistwrap cat > $@

data/liga2022/results.json: data/liga2022/description.json $(tttscraper_bin)
	./$(tttscraper_bin) --description $<  --output $@

data/liga2022/data.json: data/liga2022/results.json $(scribe_bin)
	./$(scribe_bin) --year 2021 --results $< --output $@

.PHONY: player_pages
player_pages: data/liga2021/data.json data/liga2022/data.json $(players_bin)
	./$(players_bin) \
		--output=content/player \
		--leagues="data/liga2021/results.json,data/liga2022/results.json"

$(tttscraper_bin): tools/cmd/tttscraper/main.go
	cd tools/cmd/tttscraper && go build .

$(scribe_bin): tools/cmd/scribe/main.go
	cd tools/cmd/scribe/ && go build .

$(players_bin): tools/cmd/players/main.go
	cd tools/cmd/players && go build .
