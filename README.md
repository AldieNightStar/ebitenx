# Ebiten-X
## Easier way to create games
* Based on `Ebiten` lib
* Generics
* You still can use `ebiten.*` and `ebitenutils.*` etc as before. Nothing special changed

# Import
```go
import "github.com/AldieNightStar/ebitenx"
```

# API
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

func draw(api *ebitenx.DrawAPI[*CurState]) {
	api.Image() // Get's ebiten.Image

	api.DrawImage(img, x, y) // Draw image on x,y coords
	api.DrawImageScale(img, x,y,1,1) // Draw image with scale multiplier

	api.SetPixel(x,y, color) // Set's pixel color at the screen
	api.DrawRect(x,y,w,h,color) // Draw Rect on the screen
	api.DrawLine(x,y,x2,y2,color) // Draw line on the screen
}

func update(api *ebitenx.GameAPI[*CurState]) error {
	api.JustPressed(ebitenx.KeyA) // bool - returns true if key is pressed JUST NOW (One time per frame)
	api.Pressed(ebitenx.KeyA) // bool - returns true if key is pressed

	api.JustMousePressed(ebitenx.MouseButtonLeft) // Mouse just pressed :: true/false
	api.JustMouseReleased(ebitenx.MouseButtonLeft) // Mouse just released :: true/false
	api.MousePressed(ebitenx.MouseButtonLeft) // Mouse pressed :: true/false
	
	api.GetMousePos() // x, y int - returns x,y mouse pos on the screen

	api.PressedKeys() // []int - Pressed Key list
    return nil
}
```