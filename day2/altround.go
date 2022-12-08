package day2

type altround struct {
	opponent choice
	target   outcome
}

func (r altround) result() int {
	return int(r.choice()) + int(r.target)
}

func (r altround) choice() choice {
	switch r.target {
	case draw:
		return r.opponent
	case win:
		switch r.opponent {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		}
	case loss:
		switch r.opponent {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	}

	panic("impossible")
}
