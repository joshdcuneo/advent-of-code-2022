package day1

import (
	"fmt"
	"sort"
)

func Part1() {
	input := newInput()
	defer input.close()

	elves := input.parse()

	fmt.Printf("The elf with the highest total has %d calories", totalAndSort(elves)[0])
}

func Part2() {
	input := newInput()
	defer input.close()

	sorted := totalAndSort(input.parse())
	t := 0

	for _, n := range sorted[:3] {
		t = t + n
	}

	fmt.Printf("The top 3 elves have %d calories", t)
}

func totalAndSort(elves []elf) []int {
	totals := make([]int, len(elves))

	for _, e := range elves {
		totals = append(totals, e.CalculateTotal())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	return totals
}
