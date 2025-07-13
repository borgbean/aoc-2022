package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"slices"
)

func part1(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	rows, cols := len(lines), len(lines[0])
	visible := make([][]bool, len(lines))

	for i := 1; i < (rows - 1); i++ {
		row := lines[i]

		visible[i] = make([]bool, cols)

		maxSuffix := row[cols-1]
		for j := cols - 2; j > 0; j-- {
			if lines[i][j] > maxSuffix {
				visible[i][j] = true
			}
			maxSuffix = max(maxSuffix, lines[i][j])
		}
		maxPrefix := row[0]
		for j := 1; j < (cols - 1); j++ {
			if lines[i][j] > maxPrefix {
				visible[i][j] = true
			}
			maxPrefix = max(maxPrefix, lines[i][j])
		}
	}

	for j := 1; j < (cols - 1); j++ {
		maxSuffix := lines[rows-1][j]
		for i := (rows - 2); i > 0; i-- {
			if lines[i][j] > maxSuffix {
				visible[i][j] = true
			}
			maxSuffix = max(maxSuffix, lines[i][j])
		}
		maxPrefix := lines[0][j]
		for i := 1; i < (rows - 1); i++ {
			if lines[i][j] > maxPrefix {
				visible[i][j] = true
			}
			maxPrefix = max(maxPrefix, lines[i][j])
		}
	}

	ret := len(lines)*2 + len(lines[0])*2 - 4
	for i := 1; i < (rows - 1); i++ {
		for j := 1; j < (cols - 1); j++ {
			if visible[i][j] {
				ret++
			}
		}
	}

	return fmt.Sprint(ret)

}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	rows, cols := len(lines), len(lines[0])
	visible := make([][]int, len(lines))

	explore := func(iOff, jOff, i, j, iEnd, jEnd int) {
		stack := [][2]int{}
		idx := 0

		for !(i == iEnd && j == jEnd) {
			//TODO binary search
			for stackidx := len(stack) - 1; stackidx >= 0; stackidx-- {
				if stack[stackidx][0] < int(lines[i][j]) {
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}

			count := idx
			if len(stack) > 0 {
				count = idx - stack[len(stack)-1][1]

				if stack[len(stack)-1][0] == int(lines[i][j]) {
					stack = stack[:len(stack)-1]
				}
			}
			visible[i][j] *= count

			stack = append(stack, [2]int{int(lines[i][j]), idx})

			idx += 1
			i += iOff
			j += jOff
		}
	}

	for i := 0; i < rows; i++ {
		visible[i] = slices.Repeat([]int{1}, cols)

		explore(0, 1, i, 0, i, cols)
		explore(0, -1, i, cols-1, i, -1)
	}

	for j := 0; j < cols; j++ {
		explore(1, 0, 0, j, rows, j)
		explore(-1, 0, rows-1, j, -1, j)
	}

	best := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			best = max(best, visible[i][j])
		}
	}

	return fmt.Sprint(best)
}

func main() {
	input, err := aoc.GetInput("2022", "8")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
