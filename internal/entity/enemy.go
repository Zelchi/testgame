package entity

import "testgame/internal/components"

type Enemy struct {
	*Sprite
	Following       bool
	CanAttackPlayer bool
	Combat          *components.EnemyCombat
}
