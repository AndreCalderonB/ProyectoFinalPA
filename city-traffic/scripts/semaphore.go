package scripts

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Semaphore struct {
	game    *Game
	cars    []*Car
	timerR  time.Duration
	timerG  time.Duration
	xPos    float64
	yPos    float64
	dTime   int
	carChan chan int
	state   bool //true - verde - false rojo
	img     ebiten.Image
}

var s Semaphore

func SemInit(g *Game, r int, d int) *Semaphore {

	s := Semaphore{
		game:   g,
		timerR: 0,
		timerG: 2,
		state:  false,
		xPos:   10,
		yPos:   10,
	}
	s.carChan = make(chan int)
	carritos := make([]*Car, r)
	for i := 0; i < r; i++ {
		carritos[i] = CarInit(s.game, 1, d, (true), &s)
	}
	s.cars = carritos

	go handleLights()

	return &s
}

func handleLights() {
	for true {
		time.Sleep(s.timerR)
		s.state = true

		time.Sleep(s.timerG)
		s.state = false
	}
}

func (s *Semaphore) Update(dTime int, t time.Duration) error {
	time.Sleep(t)
	s.dTime = (s.dTime + 1) % 20
	for _, car := range s.cars {
		if err := car.Update(s.dTime); err != nil {
			s.carChan <- s.dTime
		}
	}

	return nil
}

func (s *Semaphore) Draw(screen *ebiten.Image) error {

	for _, car := range s.cars {
		if err := car.Draw(screen); err != nil {
			return err
		}
	}
	cDo := &ebiten.DrawImageOptions{}
	cDo.GeoM.Translate(s.xPos, s.yPos)
	return nil
}

func addCar(c *Car) {
	s.cars = append(s.cars, c)
}

func removeCar() *Car {
	if len(s.cars) > 0 {
		res := s.cars[1]
		s.cars = s.cars[1:]
		return res
	}
	return nil
}
