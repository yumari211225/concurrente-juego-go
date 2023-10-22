package main

import (
	"log"
	"NAVEESPACIAL2/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	gameScene := scenes.NewGameScene()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Nave Espacial")

	if err := ebiten.RunGame(gameScene); err != nil {
		log.Fatal(err)
	}
}
