package ebitenx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game[STATE any] struct {
	screenWidth  int
	screenHeight int
	State        STATE
	Updater      Updater[STATE]
	Drawer       Drawer[STATE]
	DrawApi      *DrawAPI
	GameAPI      *GameAPI[STATE]
}

func (g *Game[STATE]) Update() error {
	if g.GameAPI == nil {
		g.GameAPI = NewGameApi(g)
	}
	if g.Updater != nil {
		return g.Updater(g.GameAPI)
	}
	return nil
}

func (g *Game[STATE]) Draw(screen *ebiten.Image) {
	if g.DrawApi == nil {
		g.DrawApi = NewDrawApi(screen)
	}
	if g.Drawer != nil {
		g.Drawer(g.State, g.DrawApi)
	}
}

func (g *Game[STATE]) Layout(int, int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game[STATE]) Loop() {
	ebiten.RunGame(g)
}
