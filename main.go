package main

import (
	"GoGame/func/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Game implements ebiten.Game interface.
func main() {
	NewGame := &game.Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Game Example")
	// Call ebiten.RunGame to start your NewGame loop.
	if err := ebiten.RunGame(NewGame); err != nil {
		log.Fatal(err)
	}
}
