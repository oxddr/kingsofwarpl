tttscraper_bin = tools/cmd/tttscraper/tttscraper
players_bin = tools/cmd/players/players
ranking_bin = tools/cmd/ranking/ranking

data/ranking2021/raw/20210918_Bialystok.json: $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/4-turniej-mistrzostw-polski-kings-of-war-2021biaostocki-debiut --output $@

data/ranking2021/raw/20210124_Warszawa.json:  $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/1-turniej-ligi-kings-of-war-2021-warszawa --output $@

data/ranking2021/raw/20210227_Katowice.json: $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/ii-turniej-cyklu-mistrzostw-polski-kings-of-war-2021 --output $@

data/ranking2021/raw/20210418_Warszawa.json: $(tttscraper_bin)
	./$(tttscraper_bin) --result_url https://tabletop.to/iii-turniej-ligi-kings-of-war-polska-2021 --output $@

data/ranking2021/ranking.json: $(ranking_bin) \
data/ranking2021/raw/20210918_Bialystok.json \
data/ranking2021/raw/20210124_Warszawa.json \
data/ranking2021/raw/20210227_Katowice.json \
data/ranking2021/raw/20210418_Warszawa.json
	./$(ranking_bin) --output $@ --results_dir data/ranking2021/raw

.PHONY: player_pages
player_pages: data/ranking2021/ranking.json players
	./$(players_bin) --output content/player --ranking_path $<

$(ranking_bin): tools/cmd/ranking/main.go
	cd tools/cmd/ranking && go build .

$(players_bin): tools/cmd/players/main.go
	cd tools/cmd/players && go build .

$(tttscraper_bin): tools/cmd/tttscraper/main.go
	cd tools/cmd/tttscraper && go build .
