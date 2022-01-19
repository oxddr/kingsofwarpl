tttscraper_bin = tools/cmd/tttscraper/tttscraper
players_bin = tools/cmd/players/players
ranking_bin = tools/cmd/ranking/ranking
statistics_bin = tools/cmd/statistics/statistics

data/events.json:
	curl -L "https://docs.google.com/spreadsheets/d/e/2PACX-1vQDorbeJ3svVWrbYYvtCNBOLif-mCCtjz45ndjxUtF0Ec_o77D20E5ejQtPcvM-YguvU1wH6BTxCaoC/pub?gid=790449776&single=true&output=csv" |  mlr --csv sort -f date  | mlr --c2j --jlistwrap cat > $@

season2021_data = data/season2021/raw/20210918_Bialystok.json \
		data/season2021/raw/20210124_Warszawa.json \
		data/season2021/raw/20210227_Katowice.json \
		data/season2021/raw/20210418_Warszawa.json


data/season2021/raw/20210918_Bialystok.json: $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/4-turniej-mistrzostw-polski-kings-of-war-2021biaostocki-debiut --output $@

data/season2021/raw/20210124_Warszawa.json:  $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/1-turniej-ligi-kings-of-war-2021-warszawa --output $@

data/season2021/raw/20210227_Katowice.json: $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/ii-turniej-cyklu-mistrzostw-polski-kings-of-war-2021 --output $@

data/season2021/raw/20210418_Warszawa.json: $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/iii-turniej-ligi-kings-of-war-polska-2021 --output $@

data/season2021/ranking.json: $(ranking_bin) $(season2021_data)
	./$(ranking_bin) --output $@ --results_dir data/season2021/raw

data/season2021/factions.json: $(statistics_bin) $(season2021_data)
	./$(statistics_bin) --mode factions --results_dir data/season2021/raw --output $@

data/season2021/tournaments.json: $(statistics_bin) $(season2021_data)
	./$(statistics_bin) --mode tournaments --results_dir data/season2021/raw --output $@

.PHONY: player_pages
player_pages: data/season2021/ranking.json players
	./$(players_bin) --output content/player --ranking_path $<

$(ranking_bin): tools/cmd/ranking/main.go
	cd tools/cmd/ranking && go build .

$(players_bin): tools/cmd/players/main.go
	cd tools/cmd/players && go build .

$(tttscraper_bin): tools/cmd/tttscraper/main.go
	cd tools/cmd/tttscraper && go build .

$(statistics_bin): tools/cmd/statistics/main.go
	cd tools/cmd/statistics/ && go build .
