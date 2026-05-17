package internal

import (
	"fmt"
	"image"
	"image/color"
	"testgame/internal/entities"
	"testgame/internal/mapper"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player      *entities.Player
	Enemies     []*entities.Enemy
	Consumables []*entities.Consumable
	Camera      *Camera
	TilemapJSON *mapper.TilemapJSON
	TilemapIMG  *ebiten.Image
	Tilesets    []mapper.Tileset
}

func (game *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.Player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		game.Player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		game.Player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		game.Player.Y += 2
	}

	for _, consumable := range game.Consumables {
		if game.Player.X == consumable.X {
			game.Player.Health += consumable.Amount
			fmt.Printf("Pick up potion Health: %d\n", game.Player.Health)
		}
	}

	for _, enemy := range game.Enemies {
		if enemy.Following {
			if enemy.X < game.Player.X {
				enemy.X += 1
			}
			if enemy.X > game.Player.X {
				enemy.X -= 1
			}
			if enemy.Y < game.Player.Y {
				enemy.Y += 1
			}
			if enemy.Y > game.Player.Y {
				enemy.Y -= 1
			}
		}
	}

	game.Camera.FollowTarget(
		game.Player.X+PIXEL_SCALE/2,
		game.Player.Y+PIXEL_SCALE/2,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
	)

	game.Camera.Constrain(
		float64(game.TilemapJSON.Layers[0].Width)*PIXEL_SCALE,
		float64(game.TilemapJSON.Layers[0].Height)*PIXEL_SCALE,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
	)

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{120, 180, 255, 255})
	options := ebiten.DrawImageOptions{}

	// Desenha o mapa
	for layerIndex, layer := range game.TilemapJSON.Layers {
		for index, id := range layer.Data {
			if id == 0 {
				continue
			}

			x := index % layer.Width
			y := index / layer.Width
			x *= PIXEL_SCALE
			y *= PIXEL_SCALE

			img := game.Tilesets[layerIndex].Image(id, PIXEL_SCALE)

			options.GeoM.Translate(float64(x), float64(y))
			options.GeoM.Translate(0.0, -(float64(img.Bounds().Dy()) + PIXEL_SCALE))
			options.GeoM.Translate(game.Camera.X, game.Camera.Y)
			screen.DrawImage(img, &options)
			options.GeoM.Reset()
		}
	}

	// Desenha os inimigos.
	for _, enemy := range game.Enemies {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(enemy.X, enemy.Y)
		options.GeoM.Translate(game.Camera.X, game.Camera.Y)
		screen.DrawImage(
			enemy.Img.SubImage(
				image.Rect(0, 0, PIXEL_SCALE, PIXEL_SCALE),
			).(*ebiten.Image), &options,
		)
		options.GeoM.Reset()
	}

	// Desenha os consumíveis.
	for _, consumable := range game.Consumables {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(consumable.X, consumable.Y)
		options.GeoM.Translate(game.Camera.X, game.Camera.Y)
		screen.DrawImage(
			consumable.Img.SubImage(
				image.Rect(0, 0, PIXEL_SCALE, PIXEL_SCALE),
			).(*ebiten.Image), &options,
		)
		options.GeoM.Reset()
	}

	// Desenha o jogador.
	options.GeoM.Translate(game.Player.X, game.Player.Y)
	options.GeoM.Translate(game.Camera.X, game.Camera.Y)
	screen.DrawImage(
		game.Player.Img.SubImage(
			image.Rect(0, 0, PIXEL_SCALE, PIXEL_SCALE),
		).(*ebiten.Image), &options,
	)
	options.GeoM.Reset()
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}
