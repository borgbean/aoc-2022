package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"time"
)

var DIRS = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func part1(input string) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	rows, cols := len(lines), len(lines[0])

	var i0, j0 int
	seen := [][]bool{}

	for i := range rows {
		if len(lines[i]) < cols {
			rows -= 1
			continue
		}
		seen = append(seen, make([]bool, cols))
		for j := range cols {
			if lines[i][j] == 'S' {
				i0, j0 = i, j
			}
		}
	}

	q := [][2]int{{i0, j0}}
	q2 := [][2]int{}
	steps := 0

	seen[i0][j0] = true

	for len(q) > 0 {
		i, j := q[len(q)-1][0], q[len(q)-1][1]
		q = q[:len(q)-1]

		for _, dir := range DIRS {
			i2, j2 := i+dir[0], j+dir[1]

			if i2 < 0 || j2 < 0 || i2 >= rows || j2 >= cols || seen[i2][j2] {
				continue
			}

			if lines[i2][j2] == 'E' {
				if lines[i][j] >= 'y' {
					return fmt.Sprint(steps + 1)
				} else {
					continue
				}
			}

			diff := int(int(lines[i2][j2]) - int(lines[i][j]))

			if lines[i][j] != 'S' && diff > 1 {
				continue
			}

			seen[i2][j2] = true
			q2 = append(q2, [2]int{i2, j2})
		}

		if len(q) < 1 {
			q, q2 = q2, q
			steps++
		}
	}

	return "wrong"
}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	rows, cols := len(lines), len(lines[0])

	var i0, j0 int
	seen := [][]bool{}

	for i := range rows {
		if len(lines[i]) < cols {
			rows -= 1
			continue
		}
		seen = append(seen, make([]bool, cols))
		for j := range cols {
			if lines[i][j] == 'E' {
				i0, j0 = i, j
			}
		}
	}

	q := [][2]int{{i0, j0}}
	q2 := [][2]int{}
	steps := 0

	seen[i0][j0] = true

	for len(q) > 0 {
		i, j := q[len(q)-1][0], q[len(q)-1][1]
		q = q[:len(q)-1]

		for _, dir := range DIRS {
			i2, j2 := i+dir[0], j+dir[1]

			if i2 < 0 || j2 < 0 || i2 >= rows || j2 >= cols || seen[i2][j2] {
				continue
			}

			if lines[i][j] == 'E' {
				if lines[i2][j2] < 'y' {
					continue
				}
			}

			if (lines[i2][j2] == 'a' || lines[i2][j2] == 'E') && lines[i][j] < 'c' {
				return fmt.Sprint(steps + 1)
			}

			diff := int(int(lines[i][j]) - int(lines[i2][j2]))

			if diff > 1 {
				continue
			}

			seen[i2][j2] = true
			q2 = append(q2, [2]int{i2, j2})
		}

		if len(q) < 1 {
			q, q2 = q2, q
			steps++
		}
	}

	return "wrong"
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "12")
	if err != nil {
		log.Fatal(err)
	}

	// 	input = `Sabqponm
	// abcryxxl
	// accszExk
	// acctuvwj
	// abdefghi`

	log.Println(part1(input))
	log.Println(part2(input))

	log.Println((time.Since(start)))
}
