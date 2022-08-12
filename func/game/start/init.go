package start

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/jpeg"
)

type Init struct {
	mainSprite *ebiten.Image
	fullFace   *ebiten.Image
}

func (Init) Start() (err error, r Init) {
	// Write your start cod here.
	r.fullFace, _, err = ebitenutil.NewImageFromFile("assets/golang_gofer.jpeg")
	if err != nil {
		return errors.New("failed when initializing sprite, error details:\n\n" + err.Error()), r
	}
	r.mainSprite, _, err = ebitenutil.NewImageFromFile("assets/golang_gofer.jpeg")
	if err != nil {
		return errors.New("failed when initializing sprite, error details:\n\n" + err.Error()), r
	}
	return err, r
}
