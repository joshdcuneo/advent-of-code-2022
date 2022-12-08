package main

import (
	"os"

	"github.com/joshdcuneo/advent-of-code-2022/day1"
	"github.com/joshdcuneo/advent-of-code-2022/day2"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "1":
		if args[1] == "1" {
			day1.Part1()
		} else {
			day1.Part2()
		}
	case "2":
		if args[1] == "1" {
			day2.Part1()
		} else {
			day2.Part2()
		}
	}
}
