package ebitenx

type Updater[STATE any] func(*Game[STATE], STATE) error
type Drawer[STATE any] func(STATE, *DrawAPI)
