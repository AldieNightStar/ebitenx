package ebitenx

import "github.com/hajimehoshi/ebiten/v2"

type Game[STATE any] struct {
	screenWidth  int
	screenHeight int
	state        STATE
	Updater      Updater[STATE]
	Drawer       Drawer[STATE]
}

func (g *Game[STATE]) Update() error {
	if g.Updater != nil {
		return g.Updater(g, g.state)
	}
	return nil
}

func (g *Game[STATE]) Draw(screen *ebiten.Image) {
	if g.Drawer != nil {
		g.Drawer(g.state, screen)
	}
}

func (g *Game[STATE]) Layout(int, int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game[STATE]) Loop() {
	ebiten.RunGame(g)
}
