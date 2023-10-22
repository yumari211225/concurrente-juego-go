package models

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Spaceship struct {
	X, Y  float64
	Image *ebiten.Image
}

func NewSpaceship() *Spaceship {
	img, _, err := ebitenutil.NewImageFromFile("assets/spaceship.png")
	if err != nil {
		log.Fatalf("loading spaceship image: %v", err)
	}
	return &Spaceship{
		X:     320,
		Y:     400,
		Image: img,
	}
}