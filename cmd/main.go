package main

import (
	"log"
	"testgame/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetTPS(60)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Test Game")
	ebiten.SetVsyncEnabled(true)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetRunnableOnUnfocused(true)
	Game := game.NewGame()
	if err := ebiten.RunGame(Game); err != nil {
		log.Fatal(err)
	}
}
