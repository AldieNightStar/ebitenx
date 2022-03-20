package ebitenx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameAPI[STATE any] struct {
	Game        *Game[STATE]
	pressedKeys []int // Updated only in *Game[STATE]
}

func (g *GameAPI[STATE]) PressedKeys() []int {
	return g.pressedKeys
}

func (g *GameAPI[STATE]) JustPressed(key int) bool {
	return inpututil.IsKeyJustPressed(ebiten.Key(key))
}

func (g *GameAPI[STATE]) Pressed(key int) bool {
	ks := g.PressedKeys()
	if len(ks) < 1 {
		return false
	}
	for _, k := range ks {
		if k == key {
			return true
		}
	}
	return false
}

func (g *GameAPI[STATE]) JustMousePressed(button int) bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(button))
}

func (g *GameAPI[STATE]) JustMouseReleased(button int) bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButton(button))
}

func (g *GameAPI[STATE]) MousePressed(button int) bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton(button))
}

func (g *GameAPI[STATE]) GetMousePos() (int, int) {
	return ebiten.CursorPosition()
}

func NewGameApi[STATE any](g *Game[STATE]) *GameAPI[STATE] {
	return &GameAPI[STATE]{g, []int{}}
}
