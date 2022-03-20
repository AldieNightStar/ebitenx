package ebitenx

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

func (d *DrawAPI[STATE]) DrawRect(x, y, w, h float64, c color.Color) {
	ebitenutil.DrawRect(d.src, x, y, w, h, c)
}

func (d *DrawAPI[STATE]) DrawLine(x1, y1, x2, y2 float64, c color.Color) {
	ebitenutil.DrawLine(d.src, x1, y1, x2, y2, c)
}

func NewDrawApi[STATE any](img *ebiten.Image, state STATE) *DrawAPI[STATE] {
	return &DrawAPI[STATE]{img, state}
}
