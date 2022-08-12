package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/jpeg"
	_ "image/png"
)

type spriteImage struct {
	fullFace *ebiten.Image
	main     *ebiten.Image
}

func (Game) Init() (err error) {
	sprite.fullFace, _, err = ebitenutil.NewImageFromFile("assets/golang_gofer.jpeg")
	if err != nil {
		return errors.New("failed when initializing sprite, error details:\n\n" + err.Error())
	}
	sprite.main, _, err = ebitenutil.NewImageFromFile("assets/mainGo.png")
	if err != nil {
		return errors.New("failed when initializing sprite, error details:\n\n" + err.Error())
	}
	return nil
}
