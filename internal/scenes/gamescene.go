package scenes

import (
	"image"
	"testgame/internal/entity"
	"testgame/internal/game"
	"testgame/internal/mapper"
	"testgame/internal/spritesheet"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	loaded            bool
	player            *entity.Player
	playerSpriteSheet *spritesheet.SpriteSheet
	enemies           []*entity.Enemy
	potions           []*entity.Consumable
	tilemapJSON       *mapper.TilemapJSON
	tilesets          []mapper.Tileset
	tilemapImg        *ebiten.Image
	cam               *game.Camera
	colliders         []image.Rectangle
}

func NewGameScene() *GameScene {
	return &GameScene{
		player:            nil,
		playerSpriteSheet: nil,
		enemies:           make([]*entity.Enemy, 0),
		potions:           make([]*entity.Consumable, 0),
		tilemapJSON:       nil,
		tilesets:          nil,
		tilemapImg:        nil,
		cam:               nil,
		colliders:         make([]image.Rectangle, 0),
		loaded:            false,
	}
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	panic("unimplemented")
}

func (g *GameScene) FirstLoad() {
	panic("unimplemented")
}

func (g *GameScene) IsLoaded() bool {
	panic("unimplemented")
}

func (g *GameScene) OnEnter() {
	panic("unimplemented")
}

func (g *GameScene) OnExit() {
	panic("unimplemented")
}

func (g *GameScene) Update() SceneId {
	panic("unimplemented")
}

var _ Scene = (*GameScene)(nil)
