package entity

type Enemy struct {
	*Sprite
	Following       bool
	CanAttackPlayer bool
}
