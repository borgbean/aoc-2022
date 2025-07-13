package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func part1(input string) string {

	ret := []string{}

	for _, line := range regexp.MustCompile("(\r*\n)+").Split(input, -1) {
		needed := 4
		counts := make([]int, 256)
		distinct := 0

		for i := 0; i < len(line); i++ {
			if i >= needed {
				oldChr := line[i-needed]
				counts[oldChr] -= 1
				if counts[oldChr] == 1 {
					distinct += 1
				} else if counts[oldChr] == 0 {
					distinct -= 1
				}
			}

			counts[line[i]] += 1
			if counts[line[i]] == 1 {
				distinct += 1
			} else if counts[line[i]] == 2 {
				distinct -= 1
			}
			if distinct == needed {
				ret = append(ret, fmt.Sprint(i+1))
				break
			}
		}
	}
	return strings.Join(ret, "\n")

}

func part2(input string) string {

	ret := []string{}

	for _, line := range regexp.MustCompile("(\r*\n)+").Split(input, -1) {
		needed := 14
		counts := make([]int, 256)
		distinct := 0

		for i := 0; i < len(line); i++ {
			if i >= needed {
				oldChr := line[i-needed]
				counts[oldChr] -= 1
				if counts[oldChr] == 1 {
					distinct += 1
				} else if counts[oldChr] == 0 {
					distinct -= 1
				}
			}

			counts[line[i]] += 1
			if counts[line[i]] == 1 {
				distinct += 1
			} else if counts[line[i]] == 2 {
				distinct -= 1
			}
			if distinct == needed {
				ret = append(ret, fmt.Sprint(i+1))
				break
			}
		}
	}
	return strings.Join(ret, "\n")
}

func main() {
	input, err := aoc.GetInput("2022", "6")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
