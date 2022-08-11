package start

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Init struct {
	mainSprite *ebiten.Image
}

func (Init) Start() (err error, r Init) {
	// Write your start cod here.
	r.mainSprite, _, err = ebitenutil.NewImageFromFile("assets/mainGo.png")
	if err != nil {
		return errors.New("failed when initializing sprite" + err.Error()), r
	}
	return err, r
}
