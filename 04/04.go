package main

import (
	"aoc2022/aoc"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {
	count := 0

	for _, line := range regexp.MustCompile("(\r*\n)+").Split(input, -1) {
		if line == "" {
			continue
		}

		split := strings.Split(line, ",")

		split2, split3 := strings.Split(split[0], "-"), strings.Split(split[1], "-")

		a1, err1 := strconv.Atoi(split2[0])
		a2, err2 := strconv.Atoi(split2[1])
		b1, err3 := strconv.Atoi(split3[0])
		b2, err4 := strconv.Atoi(split3[1])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Panic("bad input " + line)
		}

		if b1 < a1 || (a1 == b1 && b2 > a2) {
			a1, a2, b1, b2 = b1, b2, a1, a2
		}

		if a1 <= b1 && a2 >= b2 {
			count += 1
		}
	}

	return count
}

func part2(input string) int {
	count := 0

	for _, line := range regexp.MustCompile("(\r*\n)+").Split(input, -1) {
		if line == "" {
			continue
		}

		split := strings.Split(line, ",")

		split2, split3 := strings.Split(split[0], "-"), strings.Split(split[1], "-")

		a1, err1 := strconv.Atoi(split2[0])
		a2, err2 := strconv.Atoi(split2[1])
		b1, err3 := strconv.Atoi(split3[0])
		b2, err4 := strconv.Atoi(split3[1])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Panic("bad input " + line)
		}

		if b1 <= a2 && b2 >= a1 {
			count += 1
		}
	}

	return count
}

func main() {
	input, err := aoc.GetInput("2022", "4")
	if err != nil {
		log.Fatal(err)
	}

	// 	input = `2-4,6-8
	// 2-3,4-5
	// 5-7,7-9
	// 2-8,3-7
	// 6-6,4-6
	// 2-6,4-8`

	log.Println(part1(input))
	log.Println(part2(input))
}
