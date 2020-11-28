package scripts

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	playing  bool
	num_cars int
	cars     []*Car
	carChan  chan int
	car      *Car
	hud      *Hud
	dTime    int
	bg       ebiten.Image
}

func NewGame(ncars int) Game {
	g := Game{playing: true, num_cars: ncars, dTime: 0}

	g.carChan = make(chan int)
	img, _, _ := ebitenutil.NewImageFromFile("imgs/bg.png", ebiten.FilterDefault)
	g.bg = *img
	rand.Seed(time.Now().Unix())
	g.car = CarInit(&g, 1, (rand.Intn(4) + 1))
	g.hud = CreateHud(&g, g.num_cars)
	return g
}

func (g *Game) Update() error {
	if g.playing {
		g.dTime = (g.dTime + 1) % 20
		if err := g.car.Update(g.dTime); err != nil {
			g.carChan <- g.dTime
		}

	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) error {
	cDo := &ebiten.DrawImageOptions{}
	cDo.GeoM.Translate(0, 0)
	screen.DrawImage(&g.bg, cDo)
	if err := g.car.Draw(screen); err != nil {
		return err
	}
	/*
		if err := g.hud.Draw(screen) {
			return err
		}
	*/
	return nil
}
