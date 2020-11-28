package main

import (
	"fmt"
	"log"

	"github.com/AndreCalderonB/City_Traffic/scripts"
	"github.com/hajimehoshi/ebiten"
)

var gm scripts.Game

func init() {
	gm = scripts.NewGame(5)
}

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	if err := gm.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if err := gm.Draw(screen); err != nil {
		fmt.Println(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

func main() {

	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Traffic Sim")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
