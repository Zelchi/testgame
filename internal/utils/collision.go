package utils

import (
	"image"
	"testgame/internal/constant"
	"testgame/internal/entity"
)

func CheckCollisionHorizontal(sprite *entity.Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(
			int(sprite.X),
			int(sprite.Y),
			int(sprite.X)+constant.TILESIZE,
			int(sprite.Y)+constant.TILESIZE,
		)) {
			if sprite.DeltaX > 0.0 {
				sprite.X = float64(collider.Min.X) - float64(constant.TILESIZE)
			}
			if sprite.DeltaX < 0.0 {
				sprite.X = float64(collider.Max.X)
			}
		}
	}
}

func CheckCollisionVertical(sprite *entity.Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(
			int(sprite.X),
			int(sprite.Y),
			int(sprite.X)+constant.TILESIZE,
			int(sprite.Y)+constant.TILESIZE,
		)) {
			if sprite.DeltaY > 0.0 {
				sprite.Y = float64(collider.Min.Y) - float64(constant.TILESIZE)
			}
			if sprite.DeltaY < 0.0 {
				sprite.Y = float64(collider.Max.Y)
			}
		}
	}
}
