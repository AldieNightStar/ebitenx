package main

import (
	"image/color"

	"github.com/AldieNightStar/ebitenx"
	"github.com/hajimehoshi/ebiten/v2"
)

type CurState struct {
	x, y int
}

func main() {
	ebitenx.NewGame(&CurState{}).Drawer(draw).Updater(update).Build().Loop()
}

func draw(state *CurState, api *ebitenx.DrawAPI) {
	api.SetPixel(state.x, state.y, color.RGBA{255, 0, 0, 0})
}

func update(api *ebitenx.GameAPI[*CurState]) error {
	if api.Pressed(ebiten.KeyLeft) {
		api.Game.State.x -= 1
	} else if api.Pressed(ebiten.KeyRight) {
		api.Game.State.x += 1
	}
	if api.Pressed(ebiten.KeyUp) {
		api.Game.State.y -= 1
	} else if api.Pressed(ebiten.KeyDown) {
		api.Game.State.y += 1
	}
	return nil
}
