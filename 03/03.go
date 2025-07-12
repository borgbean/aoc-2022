package main

import (
	"aoc2022/aoc"
	"log"
	"regexp"
)

func getPri(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return 1 + b - 'a'
	}
	return 27 + b - 'A'
}

func part1(input string) int {
	sum := 0
	for _, line := range regexp.MustCompile("\r?\n+").Split(input, -1) {
		mid := len(line) / 2

		seenL := [53]bool{}
		seenR := [53]bool{}
		for i := 0; i < mid; i++ {
			seenL[getPri(line[i])] = true
			seenR[getPri(line[i+mid])] = true
		}

		for i := range seenL {
			if seenL[i] && seenR[i] {
				sum += i

			}
		}
	}
	return sum
}

func part2(input string) int {
	sum := 0
	seen := [3][53]int{}
	idx, it := 0, 1

	for _, line := range regexp.MustCompile("\r?\n+").Split(input, -1) {
		for i := range len(line) {
			seen[idx][getPri(line[i])] = it
		}

		idx = (idx + 1) % 3
		if idx != 0 {
			continue
		}
		for i := 0; i < len(seen[0]); i++ {
			match := true
			for j := 0; j < len(seen); j++ {
				if seen[j][i] != it {
					match = false
					break
				}
			}
			if match {
				sum += i
				break
			}
		}

		it += 1
	}
	return sum
}

func main() {
	input, err := aoc.GetInput("2022", "3")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(part1(input))
	log.Println(part2(input))
}
