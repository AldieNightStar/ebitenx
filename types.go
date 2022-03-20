package ebitenx

type Updater[STATE any] func(*GameAPI[STATE]) error
type Drawer[STATE any] func(STATE, *DrawAPI)
