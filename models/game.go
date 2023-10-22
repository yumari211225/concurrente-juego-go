package models

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
    Spaceship  *Spaceship
    Asteroids  []*Asteroid
    Score      int
    Background *ebiten.Image
    GameOver   bool
}

func NewGame() *Game {
	bg, _, err := ebitenutil.NewImageFromFile("assets/fondo.png")
	if err != nil {
		log.Fatalf("loading background image: %v", err)
	}

	return &Game{
		Spaceship:  NewSpaceship(),
		Asteroids:  []*Asteroid{},
		Score:      0,
		Background: bg,
		GameOver:   false,
	}
}

func (g *Game) CollisionDetected() bool {
    for _, asteroid := range g.Asteroids {
        if g.Spaceship.X < asteroid.X+float64(asteroid.Image.Bounds().Dx()) &&
           g.Spaceship.X+float64(g.Spaceship.Image.Bounds().Dx()) > asteroid.X &&
           g.Spaceship.Y < asteroid.Y+float64(asteroid.Image.Bounds().Dy()) &&
           g.Spaceship.Y+float64(g.Spaceship.Image.Bounds().Dy()) > asteroid.Y {
            return true
        }
    }
    return false
}

func (g *Game) UpdateScore() {
    for i, asteroid := range g.Asteroids {
        if asteroid.Y > 480 {
            g.Score++
            // Remueve el asteroide que ya pasó
            g.Asteroids = append(g.Asteroids[:i], g.Asteroids[i+1:]...)
        }
    }
}
