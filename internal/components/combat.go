package components

type Combat interface {
	Health() int
	AttackPower() int
	Attacking() bool
	Attack() bool
	Damage(amount int)
}

type BasicCombat struct {
	health       int
	attack_power int
	attacking    bool
}

func NewBasicCombat(health, attack_power int) *BasicCombat {
	return &BasicCombat{
		health:       health,
		attack_power: attack_power,
		attacking:    false,
	}
}

func (combat *BasicCombat) Health() int {
	return combat.health
}

func (combat *BasicCombat) AttackPower() int {
	return combat.attack_power
}

func (combat *BasicCombat) Attacking() bool {
	return combat.attacking
}

func (combat *BasicCombat) Attack() bool {
	combat.attacking = true
	return true
}

func (combat *BasicCombat) Update() {}

func (combat *BasicCombat) Damage(amount int) {
	combat.health -= amount
}

var _ Combat = (*BasicCombat)(nil)

type EnemyCombat struct {
	*BasicCombat
	attackCooldown  int
	timeSinceAttack int
}

func NewEnemyCombat(health, attack_power, attackCooldown int) *EnemyCombat {
	return &EnemyCombat{
		BasicCombat:     NewBasicCombat(health, attack_power),
		attackCooldown:  attackCooldown,
		timeSinceAttack: attackCooldown,
	}
}

func (combat *EnemyCombat) Attack() bool {
	if combat.timeSinceAttack >= combat.attackCooldown {
		combat.attacking = true
		combat.timeSinceAttack = 0
		return true
	}
	return false
}

func (combat *EnemyCombat) Update() {
	combat.timeSinceAttack += 1
}

var _ Combat = (*EnemyCombat)(nil)
