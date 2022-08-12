package main

import (
	"GoGame/func/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Game implements ebiten.Game interface.
const (
	screenWidth  = 1600
	screenHeight = 900
)

func main() {
	NewGame := &game.Game{}
	err := NewGame.Start()
	if err != nil {
		log.Panic(err)
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Game Example")
	if err := ebiten.RunGame(NewGame); err != nil {
		log.Fatal(err)
	}
}
