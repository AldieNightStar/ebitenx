package main

import (
	"image/color"

	"github.com/AldieNightStar/ebitenx"
	"github.com/hajimehoshi/ebiten/v2"
)

type CurState struct{}

func main() {
	ebitenx.NewGame(&CurState{}).Drawer(draw).Updater(update).Build().Loop()
}

func draw(state *CurState, screen *ebiten.Image) {
	screen.Set(10, 10, color.RGBA{255, 0, 0, 0})
}

func update(game *ebitenx.Game[*CurState], state *CurState) error {
	return nil
}
