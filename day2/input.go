package day2

import (
	"bufio"
	"os"
	"strings"
)

type input struct {
	filepath string
	file     *os.File
}

func newInput() input {
	return input{filepath: "day2/input.txt", file: nil}
}

func (i *input) parse() []round {
	scanner := i.scan()

	var rounds []round

	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		rounds = append(rounds, round{opponent: parseChoice(c[0]), self: parseChoice(c[1])})
	}

	return rounds
}

func (i *input) altparse() []altround {
	scanner := i.scan()

	var rounds []altround

	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		rounds = append(rounds, altround{opponent: parseChoice(c[0]), target: parseOutcome(c[1])})
	}

	return rounds
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
