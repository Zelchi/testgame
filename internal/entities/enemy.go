package entities

type Enemy struct {
	*Sprite
	Following       bool
	CanAttackPlayer bool
}
