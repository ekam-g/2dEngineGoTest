package game

import (
	_ "image/jpeg"
	_ "image/png"
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}
