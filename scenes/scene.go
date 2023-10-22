package scenes

import (
	"fmt"
	"NAVEESPACIAL2/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
)

type GameScene struct {
	game *models.Game
}

func NewGameScene() *GameScene {
	return &GameScene{
		game: models.NewGame(),
	}
}

func (g *GameScene) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.game.Spaceship.X > 0 {
		g.game.Spaceship.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.game.Spaceship.X < 640-32 {
		g.game.Spaceship.X += 5
	}

	for _, asteroid := range g.game.Asteroids {
		asteroid.Y += asteroid.Velocity
	}

	if rand.Float64() < 0.02 {
		g.game.Asteroids = append(g.game.Asteroids, models.NewAsteroid())
	}

	if g.game.CollisionDetected() {
		g.game.GameOver = true
	}

	if !g.game.GameOver {
		g.game.UpdateScore()
	}

	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.game.Background, &ebiten.DrawImageOptions{})
	
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.game.Spaceship.X, g.game.Spaceship.Y)
	screen.DrawImage(g.game.Spaceship.Image, opts)

	for _, asteroid := range g.game.Asteroids {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(asteroid.X, asteroid.Y)
		screen.DrawImage(asteroid.Image, opts)
	}

	scoreText := fmt.Sprintf("Score: %d", g.game.Score)
	ebitenutil.DebugPrintAt(screen, scoreText, 540, 10)

	if g.game.GameOver {
		ebitenutil.DebugPrintAt(screen, "Perdiste!", 270, 240)
	}
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
