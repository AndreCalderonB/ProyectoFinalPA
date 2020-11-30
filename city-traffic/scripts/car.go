package scripts

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Car struct {
	game  *Game
	speed int
	dir   int
	ind   int
	xPos  float64
	yPos  float64
	img   ebiten.Image
	run   bool
	turn  bool
	sem   *Semaphore
}

func CarInit(g *Game, s int, d int, t bool, semaphore *Semaphore) *Car {

	c := Car{
		game:  g,
		speed: s,
		dir:   d,
		xPos:  50,
		yPos:  325,
		turn:  t,
		run:   true,
		sem:   semaphore,
	}

	switch d := c.dir; d {
	case 1: // Oeste-Este
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 0 // giro 250
		c.yPos = 325
	case 2: // Sur-Norte
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_a.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 325
		c.yPos = 600 // giro 250
	case 3: // Este-Oeste
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_i.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 600 // giro 325
		c.yPos = 250
	case 4: // Norte-Sur
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_ab.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 250
		c.yPos = 0 // giro 325
	}

	return &c
}

func (c *Car) Update(dTime int) error {
	if !c.run && c.sem.state {
		carStart(c)
	}
	if c.run {
		switch d := c.dir; d {
		case 1: // giro x = 250
			if c.xPos == 220 && !c.sem.state {
				carStop(c)
			}
			if c.xPos == 250 && c.turn {
				// turn down the car gui
				img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_ab.png", ebiten.FilterDefault)
				c.img = *img
				// change direction
				c.dir = 4
			} else {
				c.xPos = c.xPos + 1
			}

		case 2: // giro y = 250
			if c.yPos == 280 && !c.sem.state {
				carStop(c)
			}
			if c.yPos == 250 && c.turn {
				// turn down the car gui
				img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito.png", ebiten.FilterDefault)
				c.img = *img
				// change direction
				c.dir = 1
			} else {
				c.yPos = c.yPos - 1
			}

		case 3: // giro x = 325
			if c.yPos == 355 && !c.sem.state {
				carStop(c)
			}

			if c.xPos == 325 && c.turn {
				// turn down the car gui
				img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_a.png", ebiten.FilterDefault)
				c.img = *img
				// change direction
				c.dir = 2
			} else {
				c.xPos = c.xPos - 1
			}

		case 4: // giro y = 325
			if c.yPos == 295 && !c.sem.state {
				carStop(c)
			}
			if c.yPos == 325 && c.turn {
				// turn down the car gui
				img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_i.png", ebiten.FilterDefault)
				c.img = *img
				// change direction
				c.dir = 3
			} else {
				c.yPos = c.yPos + 1
			}
		}
	}
	return nil
}

func (c *Car) Draw(screen *ebiten.Image) error {

	cDo := &ebiten.DrawImageOptions{}
	cDo.GeoM.Translate(c.xPos, c.yPos)

	screen.DrawImage(&c.img, cDo)

	return nil
}

func carStop(c *Car) {
	c.run = false
}

func carStart(c *Car) {
	c.run = true
}
