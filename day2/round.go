package day2

type round struct {
	opponent choice
	self     choice
}

func (r round) result() int {
	return int(r.outcome()) + int(r.self)
}

func (r round) outcome() outcome {
	if r.opponent == r.self {
		return draw
	}

	switch r.self {
	case rock:
		if r.opponent == scissors {
			return win
		}
	case paper:
		if r.opponent == rock {
			return win
		}
	case scissors:
		if r.opponent == paper {
			return win
		}
	}

	return loss
}
