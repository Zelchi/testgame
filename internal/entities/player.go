package entities

type Player struct {
	*Sprite
	Health         uint
	Dashing        bool
	CanAttackEnemy bool
}
