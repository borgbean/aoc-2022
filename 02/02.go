package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
)

func part1(input string) int {
	score := 0
	for _, line := range regexp.MustCompile("\r*\n").Split(input, -1) {
		if len(line) < 1 {
			continue
		}
		op1 := line[0] - 'A'
		op2 := line[2] - 'X'

		score += int(op2) + 1
		if op1 == op2 {
			score += 3
		} else {
			if (op1 == 0 && op2 == 1) ||
				(op1 == 1 && op2 == 2) ||
				(op1 == 2 && op2 == 0) {
				score += 6
			}
		}
	}

	return score
}

func part2(input string) int {
	score := 0
	for _, line := range regexp.MustCompile("\r*\n").Split(input, -1) {
		if len(line) < 1 {
			continue
		}
		op1 := line[0] - 'A'
		outcome := line[2] - 'X'
		var op2 byte

		// lose, draw, win
		if outcome == 0 {
			/*
				0 -> 1
				1 -> 2
				2 -> 0
			*/
			op2 = (3 + op1 - 1) % 3
		} else if outcome == 1 {
			op2 = op1
			score += 3
		} else {
			/*
				0 -> 2
				1 -> 0
				2 -> 1
			*/
			op2 = (op1 + 1) % 3
			score += 6
		}

		score += int(op2) + 1
	}

	return score
}

func main() {
	input, err := aoc.GetInput("2022", "2")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
