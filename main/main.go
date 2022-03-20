package main

import (
	"github.com/AldieNightStar/ebitenx"
)

type CurState struct {
	x, y   int
	mx, my int
	mb     bool
}

func main() {
	ebitenx.NewGame(&CurState{10, 10, 0, 0, false}, update, draw).Loop()
}

func draw(api *ebitenx.DrawAPI[*CurState]) {
	api.DrawLine(0, 0, float64(api.State.x), float64(api.State.y), ebitenx.COLOR_BLUE)
	api.DrawRect(float64(api.State.x), float64(api.State.y), 10, 10, ebitenx.ColorOf(128, uint8(api.State.mx), uint8(api.State.my)))

	api.DrawRect(float64(api.State.mx), float64(api.State.my), 10, 10, ebitenx.COLOR_CYAN)
	if api.State.mb {
		api.DrawRect(100, 100, 25, 25, ebitenx.COLOR_DARK_RED)
	}
}

func update(api *ebitenx.GameAPI[*CurState]) error {
	if api.Pressed(ebitenx.KeyLeft) {
		api.Game.State.x -= 1
	} else if api.Pressed(ebitenx.KeyRight) {
		api.Game.State.x += 1
	}
	if api.Pressed(ebitenx.KeyUp) {
		api.Game.State.y -= 1
	} else if api.Pressed(ebitenx.KeyDown) {
		api.Game.State.y += 1
	}
	api.Game.State.mx, api.Game.State.my = api.GetMousePos()
	api.Game.State.mb = api.MousePressed(ebitenx.MouseButtonLeft)
	return nil
}
