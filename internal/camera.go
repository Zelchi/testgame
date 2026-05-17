package internal

import "math"

type Camera struct {
	X, Y float64
}

func NewCamera(X, Y float64) *Camera {
	return &Camera{X, Y}
}

func (camera *Camera) FollowTarget(targetX, targetY, screenWidth, screenHeight float64) {
	camera.X = -targetX + screenWidth/2.0
	camera.Y = -targetY + screenHeight/2.0
}

func (camera *Camera) Constrain(tilemapWidth, tilemapHeight, screenWidth, screenHeight float64) {
	camera.X = math.Min(camera.X, 0.0)
	camera.Y = math.Min(camera.Y, 0.0)

	camera.X = math.Max(camera.X, screenWidth-tilemapWidth)
	camera.Y = math.Max(camera.Y, screenHeight-tilemapHeight)
}
