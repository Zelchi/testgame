package utils

import (
	"path"
	"strings"
)

func ResolveEmbeddedPath(tilesetPath, rel string) string {
	rel = strings.ReplaceAll(rel, "\\", "/")
	rel = strings.TrimPrefix(rel, "/")
	return path.Clean(path.Join(path.Dir(tilesetPath), rel))
}
