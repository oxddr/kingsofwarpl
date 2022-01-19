tttscraper_bin = tools/cmd/tttscraper/tttscraper
players_bin = tools/cmd/players/players
ranking_bin = tools/cmd/ranking/ranking
statistics_bin = tools/cmd/statistics/statistics

data/events.json:
	curl -L "https://docs.google.com/spreadsheets/d/e/2PACX-1vQDorbeJ3svVWrbYYvtCNBOLif-mCCtjz45ndjxUtF0Ec_o77D20E5ejQtPcvM-YguvU1wH6BTxCaoC/pub?gid=790449776&single=true&output=csv" |  mlr --csv sort -f date  | mlr --c2j --jlistwrap cat > $@

data/liga2021/description.json:
	curl -L "https://docs.google.com/spreadsheets/d/e/2PACX-1vQDorbeJ3svVWrbYYvtCNBOLif-mCCtjz45ndjxUtF0Ec_o77D20E5ejQtPcvM-YguvU1wH6BTxCaoC/pub?gid=0&single=true&output=csv" |  mlr --csv sort -f date  | mlr --c2j --jlistwrap cat > $@

data/liga2021/results.json: data/liga2021/description.json $(tttscraper_bin)
	./$(tttscraper_bin) --description $<  --output $@

data/liga2021/ranking.json: data/liga2021/results.json $(ranking_bin)
	./$(ranking_bin) --results $< --output $@

data/liga2021/factions.json: data/liga2021/results.json $(statistics_bin)
	./$(statistics_bin) --mode factions --results $< --output $@

data/liga2021/tournaments.json: data/liga2021/results.json $(statistics_bin)
	./$(statistics_bin) --mode tournaments --results $< --output $@

.PHONY: player_pages
player_pages: data/liga2021/ranking.json players
	./$(players_bin) --output content/player --ranking_path $<

$(ranking_bin): tools/cmd/ranking/main.go
	cd tools/cmd/ranking && go build .

$(players_bin): tools/cmd/players/main.go
	cd tools/cmd/players && go build .

$(tttscraper_bin): tools/cmd/tttscraper/main.go
	cd tools/cmd/tttscraper && go build .

$(statistics_bin): tools/cmd/statistics/main.go
	cd tools/cmd/statistics/ && go build .
