# City-Traffic Simulator 2020 automation test

build:
	go get github.com/hajimehoshi/ebiten
	go get github.com/hajimehoshi/ebiten/ebitenutil
	go get golang.org/x/image/font/basicfont
	go get github.com/AndreCalderonB/City_Traffic/scripts
	go build main.go

test: build
	@echo Test City-Traffic - main.go
	./main