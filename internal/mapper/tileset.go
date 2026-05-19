package mapper

import (
	"encoding/json"
	"image"
	"testgame/internal/assets"
	"testgame/internal/constant"
	"testgame/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tileset interface {
	Image(id int) *ebiten.Image
}

type UniformTilesetJSON struct {
	Path string `json:"image"`
	GID  int
}

type UniformTileset struct {
	IMG *ebiten.Image
	GID int
}

func (u *UniformTileset) Image(id int) *ebiten.Image {
	id -= u.GID

	srcX := id % 22
	srcY := id / 22

	srcX *= constant.TILESIZE
	srcY *= constant.TILESIZE

	return u.IMG.SubImage(
		image.Rect(
			srcX, srcY, srcX+constant.TILESIZE, srcY+constant.TILESIZE,
		),
	).(*ebiten.Image)
}

type TileJSON struct {
	ID     int    `json:"id"`
	Path   string `json:"image"`
	Width  int    `json:"imagewidth"`
	Height int    `json:"imageheight"`
}

type DynamicTilesetJSON struct {
	Tiles []TileJSON `json:"tiles"`
	gid   int
}

type DynamicTileset struct {
	IMGS []*ebiten.Image
	GID  int
}

func (d *DynamicTileset) Image(id int) *ebiten.Image {
	id -= d.GID
	return d.IMGS[id]
}

func loadEmbeddedImage(p string) (*ebiten.Image, error) {
	f, err := assets.Files.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func NewTileset(pathToTileset string, gid int) (Tileset, error) {
	contents, err := assets.Files.ReadFile(pathToTileset)
	if err != nil {
		return nil, err
	}

	var DynamicTilesetJSON DynamicTilesetJSON
	if err := json.Unmarshal(contents, &DynamicTilesetJSON); err != nil {
		return nil, err
	}

	if len(DynamicTilesetJSON.Tiles) > 0 {
		DynamicTileset := DynamicTileset{
			GID:  gid,
			IMGS: make([]*ebiten.Image, 0, len(DynamicTilesetJSON.Tiles)),
		}

		for _, tileJSON := range DynamicTilesetJSON.Tiles {
			imgPath := utils.ResolveEmbeddedPath(pathToTileset, tileJSON.Path)
			img, err := loadEmbeddedImage(imgPath)
			if err != nil {
				return nil, err
			}
			DynamicTileset.IMGS = append(DynamicTileset.IMGS, img)
		}

		return &DynamicTileset, nil
	}

	var uniformTilesetJSON UniformTilesetJSON
	if err := json.Unmarshal(contents, &uniformTilesetJSON); err != nil {
		return nil, err
	}

	uniformTileset := UniformTileset{GID: gid}
	imgPath := utils.ResolveEmbeddedPath(pathToTileset, uniformTilesetJSON.Path)
	img, err := loadEmbeddedImage(imgPath)
	if err != nil {
		return nil, err
	}
	uniformTileset.IMG = img

	return &uniformTileset, nil
}
