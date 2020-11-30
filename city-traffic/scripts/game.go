package scripts

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	sem      *Semaphore
	playing  bool
	num_cars int
	//cars     []*Car
	carChan chan int
	//car      *Car
	hud   *Hud
	dTime int
	bg    ebiten.Image
}

func NewGame(ncars int) Game {
	g := Game{playing: true, num_cars: ncars, dTime: 0}

	g.carChan = make(chan int)
	img, _, _ := ebitenutil.NewImageFromFile("imgs/bg.png", ebiten.FilterDefault)
	g.bg = *img
	rand.Seed(time.Now().Unix())
	sema := SemInit(&g, 1, 1)
	g.sem = sema
	//g.car = CarInit(&g, 1, (rand.Intn(4) + 1), ( /*rand.Intn(2)%2 == 1*/ true), g.sem)
	g.hud = CreateHud(&g, g.num_cars)
	return g
}

func (g *Game) Update() error {
	if g.playing {
		g.dTime = (g.dTime + 1) % 20
		if err := g.sem.Update(g.dTime, 2); err != nil {
			g.carChan <- g.dTime
		}
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) error {
	cDo := &ebiten.DrawImageOptions{}
	cDo.GeoM.Translate(0, 0)
	screen.DrawImage(&g.bg, cDo)
	if err := g.sem.Draw(screen); err != nil {
		return err
	}
	/*
		if err := g.hud.Draw(screen) {
			return err
		}
	*/
	return nil
}
