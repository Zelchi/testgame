package internal

import (
	"fmt"
	"image"
	"image/color"
	"testgame/internal/entities"
	"testgame/internal/mapper"
	"testgame/internal/spritesheet"
	"testgame/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Player            *entities.Player
	PlayerSpriteSheet *spritesheet.SpriteSheet
	AnimationFrame    int
	Enemies           []*entities.Enemy
	Consumables       []*entities.Consumable
	Camera            *Camera
	TilemapJSON       *mapper.TilemapJSON
	TilemapIMG        *ebiten.Image
	Tilesets          []mapper.Tileset
	Colliders         []image.Rectangle
}

func (game *Game) Update() error {

	game.Player.DeltaX = 0
	game.Player.DeltaY = 0

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.Player.DeltaX = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		game.Player.DeltaX = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		game.Player.DeltaY = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		game.Player.DeltaY = 2
	}

	game.Player.X += game.Player.DeltaX
	utils.CheckCollisionHorizontal(game.Player.Sprite, game.Colliders)

	game.Player.Y += game.Player.DeltaY
	utils.CheckCollisionVertical(game.Player.Sprite, game.Colliders)

	for _, enemy := range game.Enemies {
		enemy.DeltaX = 0
		enemy.DeltaY = 0
		if enemy.Following {
			if enemy.X < game.Player.X {
				enemy.DeltaX = 1
			}
			if enemy.X > game.Player.X {
				enemy.DeltaX = -1
			}
			if enemy.Y < game.Player.Y {
				enemy.DeltaY = 1
			}
			if enemy.Y > game.Player.Y {
				enemy.DeltaY = -1
			}
		}

		enemy.X += enemy.DeltaX
		utils.CheckCollisionHorizontal(enemy.Sprite, game.Colliders)

		enemy.Y += enemy.DeltaY
		utils.CheckCollisionVertical(enemy.Sprite, game.Colliders)

		activeAnimation := game.Player.ActiveAnimation(game.Player.DeltaX, game.Player.DeltaY)
		if activeAnimation != nil {
			activeAnimation.Update()
		}

		if enemy.X == game.Player.X && enemy.Y == game.Player.Y {
			game.Player.Health -= 1
			fmt.Printf("Player hit! Health: %d\n", game.Player.Health)
		}
	}

	for _, consumable := range game.Consumables {
		if game.Player.X == consumable.X {
			game.Player.Health += consumable.Amount
			fmt.Printf("Pick up potion Health: %d\n", game.Player.Health)
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
	playerFrame := 0
	activeAnimation := game.Player.ActiveAnimation(game.Player.DeltaX, game.Player.DeltaY)
	if activeAnimation != nil {
		playerFrame = activeAnimation.Frame()
	}

	options.GeoM.Translate(game.Player.X, game.Player.Y)
	options.GeoM.Translate(game.Camera.X, game.Camera.Y)
	screen.DrawImage(
		game.Player.Img.SubImage(
			game.PlayerSpriteSheet.Rect(playerFrame),
		).(*ebiten.Image), &options,
	)
	options.GeoM.Reset()

	for _, collider := range game.Colliders {
		vector.StrokeRect(
			screen,
			float32(collider.Min.X)+float32(game.Camera.X),
			float32(collider.Min.Y)+float32(game.Camera.Y),
			float32(collider.Dx()),
			float32(collider.Dy()),
			float32(1),
			color.RGBA{255, 0, 0, 255},
			true,
		)
	}

}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}
