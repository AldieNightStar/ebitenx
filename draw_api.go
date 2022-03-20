package ebitenx

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawAPI[STATE any] struct {
	src   *ebiten.Image
	State STATE
}

func (d *DrawAPI[STATE]) Image() *ebiten.Image {
	return d.src
}

func (d *DrawAPI[STATE]) DrawImageScale(img *ebiten.Image, x, y, scaleX, scaleY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Scale(scaleX, scaleY)
	d.src.DrawImage(img, op)
}

func (d *DrawAPI[STATE]) DrawImage(img *ebiten.Image, x, y float64) {
	d.DrawImageScale(img, x, y, 1, 1)
}

func (d *DrawAPI[STATE]) SetPixel(x, y int, c color.Color) {
	d.src.Set(x, y, c)
}

func NewDrawApi[STATE any](img *ebiten.Image, state STATE) *DrawAPI[STATE] {
	return &DrawAPI[STATE]{img, state}
}
