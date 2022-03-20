package ebitenx

import "github.com/hajimehoshi/ebiten/v2"

type Updater[STATE any] func(*Game[STATE], STATE) error
type Drawer[STATE any] func(STATE, *ebiten.Image)
