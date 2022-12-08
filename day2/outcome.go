package day2

import "fmt"

type outcome int

const (
	loss outcome = 0
	draw outcome = 3
	win  outcome = 6
)

func parseOutcome(s string) outcome {
	switch s {
	case "X":
		return loss
	case "Y":
		return draw
	case "Z":
		return win
	}

	panic(fmt.Sprintf("invalid outcome: %s", s))
}
