package day1

import (
	"bufio"
	"os"
	"strconv"
)

type Input struct {
	filepath string
	file     *os.File
}

func NewInput() Input {
	return Input{filepath: "day1/input.txt", file: nil}
}

func (i *Input) Parse() []Elf {
	scanner := i.scan()

	var elves []Elf
	var inventory []int

	for scanner.Scan() {
		if line := parseLine(scanner.Text()); line != nil {
			inventory = append(inventory, *line)
		} else {
			elves = append(elves, NewElf(inventory))
			inventory = []int{}
		}
	}

	elves = append(elves, NewElf(inventory))

	return elves
}

func (i *Input) Close() {
	if err := i.file.Close(); err != nil {
		panic(err)
	}
}

func (i *Input) scan() *bufio.Scanner {
	i.open()

	fileScanner := bufio.NewScanner(i.file)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func (i *Input) open() {
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
