package main

import (
	"GoGame/func/game"
	"GoGame/func/game/start"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Game implements ebiten.Game interface.
const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	err, _ := start.Init{}.Start()
	if err != nil {
		panic(err)
	}
	NewGame := &game.Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Game Example")
	// Call ebiten.RunGame to start your NewGame loop.
	if err := ebiten.RunGame(NewGame); err != nil {
		log.Fatal(err)
	}
}
