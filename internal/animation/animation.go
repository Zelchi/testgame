package animation

type Animation struct {
	First   int
	Last    int
	Step    int
	Speed   float32
	Counter float32
	Current int
}

func NewAnimation(first, last, step int, speed float32) *Animation {
	return &Animation{
		first,
		last,
		step,
		speed,
		speed,
		first,
	}
}

func (animation *Animation) Frame() int {
	return animation.Current
}

func (animation *Animation) Update() {
	animation.Counter -= 1.0
	if animation.Counter <= 0.0 {
		animation.Counter = animation.Speed
		animation.Current += animation.Step
		if animation.Current > animation.Last {
			animation.Current = animation.First
		}
	}
}
