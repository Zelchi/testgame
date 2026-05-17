package main

import (
	"log"
	"testgame/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(60)
	Game := internal.LoadGame()
	if err := ebiten.RunGame(Game); err != nil {
		log.Fatal(err)
	}
}
