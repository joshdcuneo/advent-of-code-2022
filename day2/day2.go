package day2

import "fmt"

func Part1() {
	input := newInput()
	defer input.close()

	rounds := input.parse()

	var score int

	for _, round := range rounds {
		score = score + int(round.result())
	}

	fmt.Printf("My score would be %d", score)
}

func Part2() {
	input := newInput()
	defer input.close()

	rounds := input.altparse()

	var score int

	for _, round := range rounds {
		score = score + int(round.result())
	}

	fmt.Printf("My score would be %d", score)
}
