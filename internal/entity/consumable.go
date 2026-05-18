package entity

type Consumable struct {
	*Sprite
	Type   string
	Amount uint
}
