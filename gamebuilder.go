package ebitenx

type GameBuilder[STATE any] struct {
	game *Game[STATE]
}

func (g *GameBuilder[STATE]) Updater(f Updater[STATE]) *GameBuilder[STATE] {
	g.game.Updater = f
	return g
}

func (g *GameBuilder[STATE]) Drawer(f Drawer[STATE]) *GameBuilder[STATE] {
	g.game.Drawer = f
	return g
}

func (g *GameBuilder[STATE]) Screen(screenWidth int, screenHeight int) *GameBuilder[STATE] {
	g.game.screenWidth = screenWidth
	g.game.screenHeight = screenHeight
	return g
}

func (g *GameBuilder[STATE]) Build() *Game[STATE] {
	return g.game
}

func NewGame[STATE any](state STATE) *GameBuilder[STATE] {
	return &GameBuilder[STATE]{
		&Game[STATE]{640, 480, state, nil, nil, nil, nil},
	}
}
