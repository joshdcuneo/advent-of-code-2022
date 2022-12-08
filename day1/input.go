package day1

import (
	"bufio"
	"os"
	"strconv"
)

type input struct {
	filepath string
	file     *os.File
}

func newInput() input {
	return input{filepath: "day1/input.txt", file: nil}
}

func (i *input) parse() []elf {
	scanner := i.scan()

	var elves []elf
	var inventory []int

	for scanner.Scan() {
		if line := parseLine(scanner.Text()); line != nil {
			inventory = append(inventory, *line)
		} else {
			elves = append(elves, newElf(inventory))
			inventory = []int{}
		}
	}

	elves = append(elves, newElf(inventory))

	return elves
}

func (i *input) close() {
	if err := i.file.Close(); err != nil {
		panic(err)
	}
}

func (i *input) scan() *bufio.Scanner {
	i.open()

	fileScanner := bufio.NewScanner(i.file)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func (i *input) open() {
	file, err := os.Open(i.filepath)
	if err != nil {
		panic(err)
	}

	i.file = file
}

func parseLine(line string) *int {
	if line == "" {
		return nil
	}

	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return &num
}
