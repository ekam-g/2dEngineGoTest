package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) handleMovement() {
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.px += 4
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.py += 4
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.px -= 4
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.py -= 4
	}

	// +1/-1 is to stop player before it reaches the border
	if g.px >= screenWidth-padding {
		g.px = screenWidth - padding - 1
	}

	if g.px <= padding {
		g.px = padding + 1
	}

	if g.py >= screenHeight-padding {
		g.py = screenHeight - padding - 1
	}

	if g.py <= padding {
		g.py = padding + 1
	}
}
