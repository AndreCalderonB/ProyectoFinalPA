package scripts

import (
	"math/rand"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Semaphore struct {
	game        *Game
	cars        []*Car
	carsAtLight []*Car
	timerR      time.Duration
	timerG      time.Duration
	xPos        float64
	yPos        float64
	position    int
	dTime       int
	carChan     chan int
	state       bool //true - verde - false rojo
	img         ebiten.Image
	imgTurn     ebiten.Image
}

var s Semaphore

func SemInit(g *Game, r int, d int, pos int, wg *sync.WaitGroup) {

	s := Semaphore{
		game:     g,
		timerR:   7,
		timerG:   5,
		state:    true,
		position: pos,
	}

	switch p := s.position; p {
	case 0: // Norte-Sur
		s.xPos = 330
		s.yPos = 550
		imgTurn, _, _ := ebitenutil.NewImageFromFile("imgs/no_u_arriba.png", ebiten.FilterDefault)
		s.imgTurn = *imgTurn
	case 1: // Oeste-Este
		s.xPos = 550
		s.yPos = 550
		imgTurn, _, _ := ebitenutil.NewImageFromFile("imgs/no_u_derecha.png", ebiten.FilterDefault)
		s.imgTurn = *imgTurn
	case 2: //
		s.xPos = 550
		s.yPos = 310
		imgTurn, _, _ := ebitenutil.NewImageFromFile("imgs/no_u_abajo.png", ebiten.FilterDefault)
		s.imgTurn = *imgTurn
	case 3:
		s.xPos = 330
		s.yPos = 310
		imgTurn, _, _ := ebitenutil.NewImageFromFile("imgs/no_u_izquierda.png", ebiten.FilterDefault)
		s.imgTurn = *imgTurn
	}

	img, _, _ := ebitenutil.NewImageFromFile("imgs/semrojo.png", ebiten.FilterDefault)

	s.img = *img
	s.carChan = make(chan int)
	s.cars = []*Car{}
	s.carsAtLight = []*Car{}

	for i := 0; i < r; i++ {
		go s.makeCars(r, i)
	}

	g.sem[pos] = &s
	wg.Done()
}
func (s *Semaphore) makeCars(r int, i int) {

	time.Sleep(time.Duration(i*3) * time.Second)

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 4
	s.queue(CarInit(s.game, 2, s.position+1, rand.Intn(max-min)+min, s, i))

}
func (s *Semaphore) toggleLight() {
	s.state = !s.state
}

func (s *Semaphore) Update(dTime int, t time.Duration) error {

	time.Sleep(t)

	s.dTime = (s.dTime + 1) % 20
	//fmt.Println(s.state)
	if s.state == true {
		img, _, _ := ebitenutil.NewImageFromFile("imgs/semrojo.png", ebiten.FilterDefault)
		s.img = *img
	} else {
		img, _, _ := ebitenutil.NewImageFromFile("imgs/semver.png", ebiten.FilterDefault)
		s.img = *img
	}
	for i := 0; i < len(s.cars); i++ {
		if err := s.cars[i].Update(s.dTime); err != nil {
			s.carChan <- s.dTime
		}
	}

	return nil
}

func (s *Semaphore) Draw(screen *ebiten.Image) error {

	for i := 0; i < len(s.cars); i++ {
		if err := s.cars[i].Draw(screen); err != nil {
			return err
		}
	}
	cDo := &ebiten.DrawImageOptions{}
	cDa := &ebiten.DrawImageOptions{}
	cDo.GeoM.Translate(s.xPos, s.yPos)
	screen.DrawImage(&s.img, cDo)
	switch p := s.position; p {
	case 0:
		cDa.GeoM.Translate(s.xPos+20, s.yPos)
		screen.DrawImage(&s.imgTurn, cDa)
	case 1:
		cDa.GeoM.Translate(s.xPos, s.yPos-20)
		screen.DrawImage(&s.imgTurn, cDa)
	case 2:
		cDa.GeoM.Translate(s.xPos-20, s.yPos)
		screen.DrawImage(&s.imgTurn, cDa)
	case 3:
		cDa.GeoM.Translate(s.xPos, s.yPos+20)
		screen.DrawImage(&s.imgTurn, cDa)
	}
	return nil
}

func (s *Semaphore) queue(c *Car) {
	s.g.hud.totl_cars++
	s.g.hud.curr_cars++
	s.cars = append(s.cars, c)
	s.queueW(c)
}

func (s *Semaphore) queueW(c *Car) {
	s.carsAtLight = append(s.carsAtLight, c)
}

func (s *Semaphore) dequeue() *Car {
	if len(s.cars) > 0 {
		res := s.cars[0]
		s.cars = s.cars[1:]
		return res
	}
	return nil
}
func (s *Semaphore) dequeueW() {
	if len(s.carsAtLight) > 0 {
		s.carsAtLight = s.carsAtLight[1:]
	}
}
