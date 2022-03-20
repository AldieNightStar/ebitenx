package ebitenx

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawAPI struct {
	src *ebiten.Image
}

func (d *DrawAPI) Image() *ebiten.Image {
	return d.src
}

func (d *DrawAPI) DrawImageScale(img *ebiten.Image, x, y, scaleX, scaleY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Scale(scaleX, scaleY)
	d.src.DrawImage(img, op)
}

func (d *DrawAPI) DrawImage(img *ebiten.Image, x, y float64) {
	d.DrawImageScale(img, x, y, 1, 1)
}

func (d *DrawAPI) SetPixel(x, y int, c color.Color) {
	d.src.Set(x, y, c)
}

func NewDrawApi(img *ebiten.Image) *DrawAPI {
	return &DrawAPI{img}
}
