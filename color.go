package ebitenx

import "image/color"

func ColorOf(r uint8, g uint8, b uint8) color.Color {
	return &color.RGBA{r, g, b, 255}
}

var (
	COLOR_RED          = ColorOf(255, 0, 0)
	COLOR_DARK_RED     = ColorOf(139, 0, 0)
	COLOR_GREEN        = ColorOf(0, 128, 0)
	COLOR_BRIGHT_GREEN = ColorOf(0, 255, 0)
	COLOR_BLUE         = ColorOf(0, 0, 255)
	COLOR_DARK_BLUE    = ColorOf(0, 0, 139)
	COLOR_YELLOW       = ColorOf(255, 255, 0)
	COLOR_WHITE        = ColorOf(255, 255, 255)
	COLOR_BLACK        = ColorOf(0, 0, 0)
	COLOR_CYAN         = ColorOf(0, 255, 255)
	COLOR_PINK         = ColorOf(255, 192, 203)
	COLOR_VIOLET       = ColorOf(200, 80, 225)
	COLOR_DARK_GREY    = ColorOf(32, 32, 32)
	COLOR_BRIGHT_GREY  = ColorOf(120, 120, 120)
	COLOR_MAGENTA      = ColorOf(255, 0, 255)
	COLOR_ORANGE       = ColorOf(255, 165, 0)
)
