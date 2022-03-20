package main

import (
	"image/color"

	"github.com/AldieNightStar/ebitenx"
)

type CurState struct {
	x, y int
}

func main() {
	ebitenx.NewGame(&CurState{10, 10}).Drawer(draw).Updater(update).Build().Loop()
}

func draw(state *CurState, api *ebitenx.DrawAPI) {
	api.SetPixel(state.x, state.y, color.RGBA{255, 0, 0, 0})
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
	return nil
}
