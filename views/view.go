package views

import "github.com/hajimehoshi/ebiten/v2"

type View interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}
