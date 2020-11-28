package scripts

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Car struct {
	game  *Game
	speed int
	xPos  float64
	yPos  float64
	img   ebiten.Image
}

func CarInit(g *Game, s int) *Car {
	c := Car{
		game:  g,
		speed: s,
		xPos:  50,
		yPos:  325,
	}
	img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito.png", ebiten.FilterDefault)
	c.img = *img

	return &c
}
func (c *Car) Update(dTime int) error {
	c.xPos = c.xPos + 1

	return nil
}
func (c *Car) Draw(screen *ebiten.Image) error {

	cDo := &ebiten.DrawImageOptions{}
	cDo.GeoM.Translate(c.xPos, c.yPos)

	screen.DrawImage(&c.img, cDo)

	return nil
}
