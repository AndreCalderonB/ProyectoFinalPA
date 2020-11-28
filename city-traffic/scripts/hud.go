package scripts

type Hud struct {
	game *Game
	cars int
}

func CreateHud(g *Game, numcars int) *Hud {
	h := Hud{
		game: g,
		cars: numcars,
	}
	return &h
}
