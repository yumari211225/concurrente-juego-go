package models

import (
	"log"
	"math/rand"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Asteroid struct {
	X, Y     float64
	Velocity float64
	Image    *ebiten.Image
}

func NewAsteroid() *Asteroid {
	img, _, err := ebitenutil.NewImageFromFile("assets/asteroid.png")
	if err != nil {
		log.Fatalf("loading asteroid image: %v", err)
	}
	return &Asteroid{
		X:       rand.Float64() * 640,
		Y:       0,
		Velocity: 1 + rand.Float64()*3,
		Image:   img,
	}
}
