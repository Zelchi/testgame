package mapper

import (
	"encoding/json"
	"image"
	"testgame/internal/assets"
	"testgame/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tileset interface {
	Image(id, pixel_scale int) *ebiten.Image
}

type UniformTilesetJSON struct {
	Path string `json:"image"`
	GID  int
}

type UniformTileset struct {
	IMG *ebiten.Image
	GID int
}

func (u *UniformTileset) Image(id, pixel_scale int) *ebiten.Image {
	id -= u.GID

	srcX := id % 22
	srcY := id / 22

	srcX *= pixel_scale
	srcY *= pixel_scale

	return u.IMG.SubImage(
		image.Rect(
			srcX, srcY, srcX+pixel_scale, srcY+pixel_scale,
		),
	).(*ebiten.Image)
}

type TileJSON struct {
	ID     int    `json:"id"`
	Path   string `json:"image"`
	Width  int    `json:"imagewidth"`
	Height int    `json:"imageheight"`
}

type DynTilesetJSON struct {
	Tiles []TileJSON `json:"tiles"`
	gid   int
}

type DynTileset struct {
	IMGS []*ebiten.Image
	GID  int
}

func (d *DynTileset) Image(id, pixel_scale int) *ebiten.Image {
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

	var dynTilesetJSON DynTilesetJSON
	if err := json.Unmarshal(contents, &dynTilesetJSON); err != nil {
		return nil, err
	}

	if len(dynTilesetJSON.Tiles) > 0 {
		dynTileset := DynTileset{
			GID:  gid,
			IMGS: make([]*ebiten.Image, 0, len(dynTilesetJSON.Tiles)),
		}

		for _, tileJSON := range dynTilesetJSON.Tiles {
			imgPath := utils.ResolveEmbeddedPath(pathToTileset, tileJSON.Path)
			img, err := loadEmbeddedImage(imgPath)
			if err != nil {
				return nil, err
			}
			dynTileset.IMGS = append(dynTileset.IMGS, img)
		}

		return &dynTileset, nil
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
