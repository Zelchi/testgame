package spritesheet

import "image"

type SpriteSheet struct {
	TileWidth  int
	TileHeight int
	TileCount  int
}

func NewSpriteSheet(w, h, t int) *SpriteSheet {
	return &SpriteSheet{w, h, t}
}

func (sprite *SpriteSheet) Rect(index int) image.Rectangle {
	x := (index % sprite.TileWidth) * sprite.TileCount
	y := (index / sprite.TileWidth) * sprite.TileCount

	return image.Rect(x, y, x+sprite.TileCount, y+sprite.TileCount)
}
