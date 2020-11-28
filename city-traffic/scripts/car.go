package scripts

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Car struct {
	game  *Game
	speed int
	dir   int
	xPos  float64
	yPos  float64
	img   ebiten.Image
}

func CarInit(g *Game, s int, d int) *Car {
	c := Car{
		game:  g,
		speed: s,
		dir:   d,
		xPos:  50,
		yPos:  325,
	}
	switch d := c.dir; d {
	case 1:
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 0
		c.yPos = 325
	case 2:
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_a.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 325
		c.yPos = 600
	case 3:
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_i.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 600
		c.yPos = 250
	case 4:
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_ab.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 250
		c.yPos = 0
	}

	return &c
}
func (c *Car) Update(dTime int) error {
	switch d := c.dir; d {
	case 1:
		c.xPos = c.xPos + 1
	case 2:
		c.yPos = c.yPos - 1
	case 3:
		c.xPos = c.xPos - 1
	case 4:
		c.yPos = c.yPos + 1
	}

	return nil
}
func (c *Car) Draw(screen *ebiten.Image) error {

	cDo := &ebiten.DrawImageOptions{}
	cDo.GeoM.Translate(c.xPos, c.yPos)

	screen.DrawImage(&c.img, cDo)

	return nil
}
