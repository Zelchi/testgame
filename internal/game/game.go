package game

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"testgame/internal/constant"
	"testgame/internal/entity"
	"testgame/internal/mapper"
	"testgame/internal/spritesheet"
	"testgame/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Player            *entity.Player
	PlayerSpriteSheet *spritesheet.SpriteSheet
	AnimationFrame    int
	Enemies           []*entity.Enemy
	Consumables       []*entity.Consumable
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
				enemy.DeltaX = 0.25
			}
			if enemy.X > game.Player.X {
				enemy.DeltaX = -0.25
			}
			if enemy.Y < game.Player.Y {
				enemy.DeltaY = 0.25
			}
			if enemy.Y > game.Player.Y {
				enemy.DeltaY = -0.25
			}
		}

		enemy.X += enemy.DeltaX
		utils.CheckCollisionHorizontal(enemy.Sprite, game.Colliders)

		enemy.Y += enemy.DeltaY
		utils.CheckCollisionVertical(enemy.Sprite, game.Colliders)

		if enemy.X == game.Player.X && enemy.Y == game.Player.Y {
			game.Player.Health -= 1
			fmt.Printf("Player hit! Health: %d\n", game.Player.Health)
		}
	}

	activeAnimation := game.Player.ActiveAnimation(game.Player.DeltaX, game.Player.DeltaY)
	if activeAnimation != nil {
		activeAnimation.Update()
	}

	clicked := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	cX, cY := ebiten.CursorPosition()
	cX -= int(game.Camera.X)
	cY -= int(game.Camera.Y)
	game.Player.Combat.Update()
	pRect := image.Rect(
		int(game.Player.X),
		int(game.Player.Y),
		int(game.Player.X)+constant.TILESIZE,
		int(game.Player.Y)+constant.TILESIZE,
	)

	deadEnemies := make(map[int]struct{})
	for index, enemy := range game.Enemies {
		enemy.Combat.Update()
		rect := image.Rect(
			int(enemy.X),
			int(enemy.Y),
			int(enemy.X)+constant.TILESIZE,
			int(enemy.Y)+constant.TILESIZE,
		)

		if rect.Overlaps(pRect) {
			if enemy.Combat.Attack() {
				game.Player.Combat.Damage(enemy.Combat.AttackPower())
				fmt.Println(
					fmt.Sprintf("player damaged. health: %d\n", game.Player.Combat.Health()),
				)
				if game.Player.Combat.Health() <= 0 {
					fmt.Println("player has died!")
				}
			}
		}

		if cX > rect.Min.X && cX < rect.Max.X && cY > rect.Min.Y && cY < rect.Max.Y {
			if clicked &&
				math.Sqrt(
					math.Pow(
						float64(cX)-game.Player.X+(constant.TILESIZE/2),
						2,
					)+math.Pow(
						float64(cY)-game.Player.Y+(constant.TILESIZE/2),
						2,
					),
				) < constant.TILESIZE*5 {
				fmt.Println("damaging enemy")
				enemy.Combat.Damage(game.Player.Combat.AttackPower())

				if enemy.Combat.Health() <= 0 {
					deadEnemies[index] = struct{}{}
					fmt.Println("enemy has been eliminated")
				}
			}
		}
	}

	if len(deadEnemies) > 0 {
		newEnemies := make([]*entity.Enemy, 0)
		for index, enemy := range game.Enemies {
			if _, exists := deadEnemies[index]; !exists {
				newEnemies = append(newEnemies, enemy)
			}
		}
		game.Enemies = newEnemies
	}

	for _, consumable := range game.Consumables {
		if game.Player.X == consumable.X {
			game.Player.Health += consumable.Amount
			fmt.Printf("Pick up potion Health: %d\n", game.Player.Health)
		}
	}

	game.Camera.FollowTarget(
		game.Player.X+constant.TILESIZE/2,
		game.Player.Y+constant.TILESIZE/2,
		constant.WINDOW_WIDTH,
		constant.WINDOW_HEIGHT,
	)

	game.Camera.Constrain(
		float64(game.TilemapJSON.Layers[0].Width)*constant.TILESIZE,
		float64(game.TilemapJSON.Layers[0].Height)*constant.TILESIZE,
		constant.WINDOW_WIDTH,
		constant.WINDOW_HEIGHT,
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
			x *= constant.TILESIZE
			y *= constant.TILESIZE

			img := game.Tilesets[layerIndex].Image(id)

			options.GeoM.Translate(float64(x), float64(y))
			options.GeoM.Translate(0.0, -(float64(img.Bounds().Dy()) + constant.TILESIZE))
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
			enemy.Texture.SubImage(
				image.Rect(0, 0, constant.TILESIZE, constant.TILESIZE),
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
			consumable.Texture.SubImage(
				image.Rect(0, 0, constant.TILESIZE, constant.TILESIZE),
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
		game.Player.Texture.SubImage(
			game.PlayerSpriteSheet.Rect(playerFrame),
		).(*ebiten.Image), &options,
	)
	options.GeoM.Reset()

	// Desenhar o cursor do mouse
	cursorX, cursorY := ebiten.CursorPosition()
	vector.StrokeRect(
		screen,
		float32(cursorX),
		float32(cursorY),
		float32(10),
		float32(10),
		float32(1),
		color.RGBA{255, 0, 0, 255},
		true,
	)

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

	options.GeoM.Reset()
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.WINDOW_WIDTH, constant.WINDOW_HEIGHT
}
