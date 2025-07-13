package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func part1(input string) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	grid := make([][]int, 800)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 200)
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		steps := strings.Split(line, " -> ")
		var i, j int
		for stepIdx, step := range steps {
			stepSplit := strings.Split(step, ",")
			i2, err := strconv.Atoi(stepSplit[0])
			if err != nil {
				log.Fatal("Bad row in step", step, err)
			}
			j2, err := strconv.Atoi(stepSplit[1])
			if err != nil {
				log.Fatal("Bad col in step", step, err)
			}

			if stepIdx == 0 {
				i, j = i2, j2
				continue
			}

			block(grid, i, j, i2, j2)

			i, j = i2, j2
		}
	}

	ret := 0

	for {
		i, j := 500, 0
		for {

			if (j + 1) == len(grid[0]) {
				grid[i][j] = 1
				return fmt.Sprint(ret)
			}

			if grid[i][j+1] == 0 {
				j += 1
				continue
			}

			if grid[i-1][j+1] == 0 {
				i -= 1
				j += 1
				continue
			}
			if grid[i+1][j+1] == 0 {
				i += 1
				j += 1
				continue
			}
			ret += 1
			grid[i][j] = 1
			break
		}
	}

	return ""
}

func direction(i, j, i2, j2 int) (int, int) {
	if i2 < i {
		return -1, 0
	}
	if i2 > i {
		return 1, 0
	}

	if j2 < j {
		return 0, -1
	}
	return 0, 1
}

func block(grid [][]int, i, j, i2, j2 int) {
	iOff, jOff := direction(i, j, i2, j2)

	for i != i2 || j != j2 {
		grid[i][j] = 1

		i += iOff
		j += jOff
	}
	grid[i][j] = 1
}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	grid := make([][]int, 800)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 200)
	}

	maxJ := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		steps := strings.Split(line, " -> ")
		var i, j int
		for stepIdx, step := range steps {
			stepSplit := strings.Split(step, ",")
			i2, err := strconv.Atoi(stepSplit[0])
			if err != nil {
				log.Fatal("Bad row in step", step, err)
			}
			j2, err := strconv.Atoi(stepSplit[1])
			if err != nil {
				log.Fatal("Bad col in step", step, err)
			}

			maxJ = max(maxJ, j)
			if stepIdx == 0 {
				i, j = i2, j2
				continue
			}

			block(grid, i, j, i2, j2)

			i, j = i2, j2
		}
	}

	for i := range len(grid) {
		grid[i][maxJ+2] = 1
	}

	ret := 0

	for {
		i, j := 500, 0
		for {

			if (j + 1) == len(grid[0]) {
				log.Panic("Not possible")
				grid[i][j] = 1
				return fmt.Sprint(ret)
			}

			if grid[i][j+1] == 0 {
				j += 1
				continue
			}

			if grid[i-1][j+1] == 0 {
				i -= 1
				j += 1
				continue
			}
			if grid[i+1][j+1] == 0 {
				i += 1
				j += 1
				continue
			}

			if i == 500 && j == 0 {
				return fmt.Sprint(ret + 1)
			}

			ret += 1
			grid[i][j] = 1
			break
		}
	}

	return ""
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "14")
	if err != nil {
		log.Fatal(err)
	}

	// 	input = `498,4 -> 498,6 -> 496,6
	// 503,4 -> 502,4 -> 502,9 -> 494,9`

	fmt.Println(part1(input))
	fmt.Println(part2(input))

	log.Println((time.Since(start)))
}
