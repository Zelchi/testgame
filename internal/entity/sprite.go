package entity

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Img            *ebiten.Image
	Scale          float64
	X, Y           float64
	DeltaX, DeltaY float64
}
