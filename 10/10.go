package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	cycle := 0
	x := 1

	ret := 0

	check := func() {
		if cycle < 221 && ((cycle-20)%40) == 0 {
			ret += x * cycle
		}
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line == "noop" {
			cycle++
			check()
		} else {
			val, err := strconv.Atoi(line[5:])
			if err != nil {
				log.Fatal("Bad input", line, err)
			}

			cycle++
			check()

			cycle++
			check()
			x += val
		}
	}

	return fmt.Sprint(ret)

}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	cycle := 0
	x := 1

	ret := []string{}
	line := [40]byte{}

	check := func() {
		idx := (cycle - 1) % 40
		line[idx] = '.'
		if max(x-idx, idx-x) < 2 {
			line[idx] = '#'
		}

		if (cycle % 40) == 0 {
			ret = append(ret, string(line[:]))
		}
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line == "noop" {
			cycle++
			check()
		} else {
			val, err := strconv.Atoi(line[5:])
			if err != nil {
				log.Fatal("Bad input", line, err)
			}

			cycle++
			check()

			cycle++
			check()
			x += val
		}
	}

	return fmt.Sprint(strings.Join(ret, "\n"))
}

func main() {
	input, err := aoc.GetInput("2022", "10")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	log.Printf("\n%s\n", part2(input))
}
