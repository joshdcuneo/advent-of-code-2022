package day1

type Elf struct {
	inventory []int
}

func NewElf(inventory []int) Elf {
	return Elf{inventory: inventory}
}

func (e Elf) CalculateTotal() int {
	total := 0

	for _, item := range e.inventory {
		total += item
	}

	return total
}
