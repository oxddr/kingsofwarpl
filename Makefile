data/ranking2021/raw/20210918_Bialystok.json: tttscraper
	./tttscraper --result_url https://tabletop.to/4-turniej-mistrzostw-polski-kings-of-war-2021biaostocki-debiut --output $@

data/ranking2021/raw/20210124_Warszawa.json: tttscraper
	./tttscraper --result_url https://tabletop.to/1-turniej-ligi-kings-of-war-2021-warszawa --output $@

data/ranking2021/raw/20210227_Katowice.json: tttscraper
	./tttscraper --result_url https://tabletop.to/ii-turniej-cyklu-mistrzostw-polski-kings-of-war-2021 --output $@

data/ranking2021/raw/20210418_Warszawa.json: tttscraper
	./tttscraper --result_url https://tabletop.to/iii-turniej-ligi-kings-of-war-polska-2021 --output $@

data/ranking2021/ranking.json: ranking \
data/ranking2021/raw/20210918_Bialystok.json \
data/ranking2021/raw/20210124_Warszawa.json \
data/ranking2021/raw/20210227_Katowice.json \
data/ranking2021/raw/20210418_Warszawa.json
	./ranking --output $@ --results_dir data/ranking2021/raw

.PHONY: player_pages
player_pages: data/ranking2021/ranking.json players
	./players --output content/player --ranking_path $<

ranking: tools/ranking/ranking.go
	cd tools/ranking && go build . && cp ranking ../../

players: tools/players/generator.go
	cd tools/players && go build . && cp players ../../

tttscraper: tools/tttscraper/scraper.go
	cd tools/tttscraper && go build . && cp tttscraper ../../
