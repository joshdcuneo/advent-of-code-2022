package day1

type elf struct {
	inventory []int
}

func newElf(inventory []int) elf {
	return elf{inventory: inventory}
}

func (e elf) CalculateTotal() int {
	total := 0

	for _, item := range e.inventory {
		total += item
	}

	return total
}
