package utils

import (
	"image"
	"testgame/internal/entities"
)

func CheckCollisionHorizontal(sprite *entities.Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(
			int(sprite.X),
			int(sprite.Y),
			int(sprite.X)+int(sprite.Scale),
			int(sprite.Y)+int(sprite.Scale),
		)) {
			if sprite.DeltaX > 0.0 {
				sprite.X = float64(collider.Min.X) - float64(sprite.Scale)
			} else if sprite.DeltaX < 0.0 {
				sprite.X = float64(collider.Max.X)
			}
		}
	}
}

func CheckCollisionVertical(sprite *entities.Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(
			int(sprite.X),
			int(sprite.Y),
			int(sprite.X)+int(sprite.Scale),
			int(sprite.Y)+int(sprite.Scale),
		)) {
			if sprite.DeltaY > 0.0 {
				sprite.Y = float64(collider.Min.Y) - float64(sprite.Scale)
			} else if sprite.DeltaY < 0.0 {
				sprite.Y = float64(collider.Max.Y)
			}
		}
	}
}
