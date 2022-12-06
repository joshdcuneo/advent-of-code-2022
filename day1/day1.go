package day1

import (
	"fmt"
	"sort"
)

func Part1() {
	input := NewInput()

	elves := input.Parse()

	fmt.Printf("The elf with the highest total has %d calories", totalAndSort(elves)[0])

	input.Close()
}

func Part2() {
	input := NewInput()

	sorted := totalAndSort(input.Parse())
	t := 0

	for _, n := range sorted[:3] {
		t = t + n
	}

	fmt.Printf("The top 3 elves have %d calories", t)

	input.Close()
}

func totalAndSort(elves []Elf) []int {
	totals := make([]int, len(elves))

	for _, e := range elves {
		totals = append(totals, e.CalculateTotal())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	return totals
}
