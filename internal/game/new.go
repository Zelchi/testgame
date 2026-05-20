package game

import (
	"image"
	_ "image/png"
	"log"
	"testgame/internal/animation"
	"testgame/internal/assets"
	"testgame/internal/components"
	"testgame/internal/entity"
	"testgame/internal/mapper"
	"testgame/internal/spritesheet"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadJSON(path string) *mapper.TilemapJSON {
	data, err := assets.Files.ReadFile(path)
	if err != nil {
		panic(err)
	}
	tilemapJSON, err := mapper.NewTilemapJSON(data)
	if err != nil {
		panic(err)
	}
	return tilemapJSON
}

func loadIMAGE(path string) *ebiten.Image {
	f, err := assets.Files.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func NewGame() *Game {
	return &Game{
		Camera:      NewCamera(0, 0),
		TilemapJSON: loadJSON("maps/map1.json"),
		TilemapIMG:  loadIMAGE("images/tileset_floor.png"),
		Tilesets:    loadJSON("maps/map1.json").GenTilesets(),
		Player: &entity.Player{
			Sprite: &entity.Sprite{
				Texture: loadIMAGE("images/darkninja.png"),
				X:       500,
				Y:       500,
			},
			Health: 10,
			Animations: map[entity.PlayerState]*animation.Animation{
				entity.Down:  animation.NewAnimation(4, 12, 4, 20),
				entity.Up:    animation.NewAnimation(5, 13, 4, 20),
				entity.Left:  animation.NewAnimation(6, 14, 4, 20),
				entity.Right: animation.NewAnimation(7, 15, 4, 20),
			},
			Combat: components.NewBasicCombat(3, 1),
		},
		PlayerSpriteSheet: spritesheet.NewSpriteSheet(4, 7, 16),
		Consumables: []*entity.Consumable{
			{
				Sprite: &entity.Sprite{
					Texture: loadIMAGE("images/heart.png"),
					X:       60,
					Y:       60,
				},
				Type:   "health",
				Amount: 20,
			},
		},
		Enemies: []*entity.Enemy{
			{
				Sprite: &entity.Sprite{
					Texture: loadIMAGE("images/skeleton.png"),
					X:       10,
					Y:       10,
				},
				Following: true,
				Combat:    components.NewEnemyCombat(3, 1, 30),
			},
			{
				Sprite: &entity.Sprite{
					Texture: loadIMAGE("images/skeleton.png"),
					X:       30,
					Y:       30,
				},
				Combat: components.NewEnemyCombat(3, 1, 30),
			},
			{
				Sprite: &entity.Sprite{
					Texture: loadIMAGE("images/skeleton.png"),
					X:       70,
					Y:       70,
				},
				Combat: components.NewEnemyCombat(3, 1, 30),
			},
		},
		Colliders: []image.Rectangle{
			image.Rect(100, 100, 116, 116),
		},
	}
}
