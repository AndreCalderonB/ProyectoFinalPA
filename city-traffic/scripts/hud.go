package scripts

import (
	"image/color"
	"strconv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
)

type Hud struct {
	g         *Game
	curr_cars int
	totl_cars int
	time      int
}

func CreateHud(game *Game) *Hud {
	h := Hud{
		g:         game,
		curr_cars: 0,
		totl_cars: 0,
	}
	return &h
}

func (h *Hud) Draw(screen *ebiten.Image) error {
	text.Draw(screen, "Spawned cars: "+strconv.Itoa(h.totl_cars), basicfont.Face7x13, 20, 20, color.Black)
	text.Draw(screen, "Finished cars: "+strconv.Itoa(h.curr_cars), basicfont.Face7x13, 20, 60, color.Black)
	text.Draw(screen, "Active semaphore: "+strconv.Itoa(h.g.semactual), basicfont.Face7x13, 20, 100, color.Black)
	return nil
}
