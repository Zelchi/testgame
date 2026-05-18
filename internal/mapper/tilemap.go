package mapper

import (
	"encoding/json"
	"path"
)

type TilemapLayerJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type TilesetLayerRef struct {
	FirstGID int    `json:"firstgid"`
	Source   string `json:"source"`
}

type TilemapJSON struct {
	Layers   []TilemapLayerJSON `json:"layers"`
	Tilesets []TilesetLayerRef  `json:"tilesets"`
}

func NewTilemapJSON(contents []byte) (*TilemapJSON, error) {
	var tilemapJSON TilemapJSON
	if err := json.Unmarshal(contents, &tilemapJSON); err != nil {
		return nil, err
	}
	return &tilemapJSON, nil
}

func (t *TilemapJSON) GenTilesets() []Tileset {
	tilesets := make([]Tileset, 0, len(t.Tilesets))

	for _, tilesetData := range t.Tilesets {
		tilesetPath := path.Join("maps", tilesetData.Source)
		tileset, err := NewTileset(tilesetPath, tilesetData.FirstGID)
		if err != nil {
			panic(err)
		}
		tilesets = append(tilesets, tileset)
	}

	return tilesets
}
