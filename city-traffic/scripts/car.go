package scripts

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Car struct {
	game   *Game
	dir    int
	des    int
	trn    int
	speed  float64
	dis    float64
	xPos   float64
	yPos   float64
	run    bool
	turned bool
	light  bool
	pass   bool
	img    ebiten.Image
	sem    *Semaphore
}

func CarInit(g *Game, spd float64, d int, ds int, s *Semaphore, i int) *Car {

	c := Car{
		game:   g,
		speed:  spd,
		dir:    d,
		dis:    0,
		xPos:   50,
		yPos:   325,
		run:    true,
		sem:    s,
		turned: false,
		trn:    ds,
		pass:   false,
	}
	c.des = (d + ds) % 4
	c.light = c.sem.state
	switch d := c.dir; d {
	case 1: // Oeste-Este
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 0 // giro 250
		c.yPos = 480
	case 2: // Sur-Norte
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_a.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 480
		c.yPos = 870 // giro 250
	case 3: // Este-Oeste
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_i.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 860 // giro 325
		c.yPos = 400
	case 4: // Norte-Sur
		img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_ab.png", ebiten.FilterDefault)
		c.img = *img
		c.xPos = 400
		c.yPos = -5 // giro 325
	}

	return &c
}

func (c *Car) Update(dTime int) error {
	// Checar semáforo
	c.light = c.sem.state
	if !c.light {
		if !c.run {
			c.carStart()
		}
	} else {
		if c.run {
			c.carStop()
		}
	}
	//Dar Vueltas
	if c.run {

		if c.dis >= 290 {
			if !c.pass {
				c.sem.dequeueW()
			}
			c.pass = true
		}
		if c.dis >= 910 {
			c.sem.dequeue()
		}

		switch t := c.trn + 1; t {
		case 1: // girar izquierda
			if c.dis >= 470 && !c.turned {
				c.dir = c.des + 1
			}
		case 2:
			//Seguir derecho
		case 3: // girar derecha
			if c.dis >= 410 && !c.turned {
				c.dir = c.des + 1
			}
		}

		switch d := c.dir; d {
		case 1:
			img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito.png", ebiten.FilterDefault)
			c.img = *img
			c.xPos += c.speed
			c.dis += c.speed
		case 2:
			img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_a.png", ebiten.FilterDefault)
			c.img = *img
			c.yPos -= c.speed
			c.dis += c.speed
		case 3:
			img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_i.png", ebiten.FilterDefault)
			c.img = *img
			c.xPos -= c.speed
			c.dis += c.speed
		case 4:
			img, _, _ := ebitenutil.NewImageFromFile("imgs/carrito_ab.png", ebiten.FilterDefault)
			c.img = *img
			c.yPos += c.speed
			c.dis += c.speed
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

func (c *Car) queuePos() float64 {
	for i := 0; i < len(c.sem.carsAtLight); i++ {
		if c.sem.carsAtLight[i] == c {
			return float64(i)
		}
	}
	return -1
}
func (c *Car) atPos() bool {
	pos := c.queuePos()
	// comparar distancia recorrida, contra distancia de semaforo segun pos
	// Ej. Posición 2 tiene que estar a = distancia de semaforo - distancia de un carro * pos
	if c.dis < (290 - 75*pos) {
		return false
	} else if c.dis == (290 - 75*pos) {
		return true
	} else {
		return true
	}
}

func (c *Car) carStop() {
	if !c.atPos() || c.pass {
		c.run = true
	} else {
		c.run = false
	}
}

func (c *Car) carStart() {
	c.run = true
}
