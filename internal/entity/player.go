package entity

import (
	"testgame/internal/animation"
	"testgame/internal/components"
)

type PlayerState uint8

const (
	Down PlayerState = iota
	Up
	Left
	Right
)

type Player struct {
	*Sprite
	Health         uint
	Dashing        bool
	CanAttackEnemy bool
	Animations     map[PlayerState]*animation.Animation
	Combat         *components.BasicCombat
}

func (player *Player) ActiveAnimation(deltaX, deltaY float64) *animation.Animation {
	if deltaX > 0 {
		return player.Animations[Right]
	}
	if deltaX < 0 {
		return player.Animations[Left]
	}
	if deltaY > 0 {
		return player.Animations[Down]
	}
	if deltaY < 0 {
		return player.Animations[Up]
	}
	return nil
}
