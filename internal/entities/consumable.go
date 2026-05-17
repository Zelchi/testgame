package entities

type Consumable struct {
	*Sprite
	Type   string
	Amount uint
}
