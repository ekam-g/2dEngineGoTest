package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/jpeg"
	_ "image/png"
)

var fullFace *ebiten.Image
var mainSprite *ebiten.Image

type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(mainSprite, nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}

func (Game) Start() (err error) {
	// Write your start cod here.
	fullFace, _, err = ebitenutil.NewImageFromFile("assets/golang_gofer.jpeg")
	if err != nil {
		return errors.New("failed when initializing sprite, error details:\n\n" + err.Error())
	}
	mainSprite, _, err = ebitenutil.NewImageFromFile("assets/mainGo.png")
	if err != nil {
		return errors.New("failed when initializing sprite, error details:\n\n" + err.Error())
	}
	return nil
}
