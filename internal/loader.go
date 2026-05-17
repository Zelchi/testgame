package internal

import (
	"image"
	_ "image/png"
	"log"
	"testgame/internal/assets"
	"testgame/internal/entities"
	"testgame/internal/mapper"

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

func LoadGame() *Game {

	tilesets, err := loadJSON("maps/map1.json").GenTilesets()
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		Camera:      NewCamera(0, 0),
		TilemapJSON: loadJSON("maps/map1.json"),
		TilemapIMG:  loadIMAGE("images/tileset_floor.png"),
		Tilesets:    tilesets,
		Player: &entities.Player{
			Sprite: &entities.Sprite{
				Img: loadIMAGE("images/darkninja.png"),
				X:   500,
				Y:   500,
			},
			Health: 10,
		},
		Consumables: []*entities.Consumable{
			{
				Sprite: &entities.Sprite{
					Img: loadIMAGE("images/heart.png"),
					X:   60,
					Y:   60,
				},
				Type:   "health",
				Amount: 20,
			},
		},
		Enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Img: loadIMAGE("images/skeleton.png"),
					X:   10,
					Y:   10,
				},
				Following: true,
			},
			{
				Sprite: &entities.Sprite{
					Img: loadIMAGE("images/skeleton.png"),
					X:   30,
					Y:   30,
				},
			},
			{
				Sprite: &entities.Sprite{
					Img: loadIMAGE("images/skeleton.png"),
					X:   70,
					Y:   70,
				},
			},
		},
	}
}
