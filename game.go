package ebitenx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game[STATE any] struct {
	ScreenWidth  int
	ScreenHeight int
	State        STATE
	Updater      Updater[STATE]
	Drawer       Drawer[STATE]
	DrawApi      *DrawAPI[STATE]
	GameAPI      *GameAPI[STATE]
}

func (g *Game[STATE]) Update() error {
	if g.GameAPI == nil {
		g.GameAPI = NewGameApi(g)
	}
	if g.Updater != nil {
		g.GameAPI.pressedKeys = convertKeysToInts(inpututil.PressedKeys())
		return g.Updater(g.GameAPI)
	}
	return nil
}

func (g *Game[STATE]) Draw(screen *ebiten.Image) {
	if g.DrawApi == nil {
		g.DrawApi = NewDrawApi(screen, g.State)
	}
	if g.Drawer != nil {
		g.Drawer(g.DrawApi)
	}
}

func (g *Game[STATE]) Layout(int, int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game[STATE]) Screen(width, height int) *Game[STATE] {
	g.ScreenWidth = width
	g.ScreenHeight = height
	return g
}

func (g *Game[STATE]) Loop() {
	ebiten.RunGame(g)
}

func NewGame[STATE any](state STATE, updater Updater[STATE], drawer Drawer[STATE]) *Game[STATE] {
	return &Game[STATE]{640, 480, state, updater, drawer, nil, nil}
}
