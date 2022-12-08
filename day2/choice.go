package day2

import "fmt"

type choice int

const (
	rock     choice = 1
	paper    choice = 2
	scissors choice = 3
)

func parseChoice(s string) choice {
	switch s {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	}

	panic(fmt.Sprintf("invalid choice: %s", s))
}
