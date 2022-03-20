# Ebiten-X
## Easier way to create games
Based on `Ebiten` lib

# Import
```go
import "github.com/AldieNightStar/ebitenx"
```

# New Game
```go
// This is state object
// Used for current game logic
type CurState struct {
	x, y   int
	mx, my int
}

// In such way we create game
// .Loop() means forever loop
ebitenx.NewGame(&CurState{10, 10, 0, 0}, update, draw).Loop()

func draw(api *ebitenx.DrawAPI[*CurState]) {}
func update(api *ebitenx.GameAPI[*CurState]) error {
    return nil
}
```