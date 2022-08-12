package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/jpeg"
	_ "image/png"
)

func (Game) Init() (err error) {
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
