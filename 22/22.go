package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"strings"
	"time"
)

type op struct {
	val        int
	v1, v2, op string
}

var DIRECTIONS = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rows := len(lines) - 2
	cols := len(lines[0])

	// grid := make([][]byte, rows)
	// for r := range rows {
	// 	grid[r] = make([]byte, cols)
	// 	for c := range len(grid[r]) {
	// 		if lines[r][c] == '.' {
	// 			grid[r][c] = 1
	// 		} else if lines[r][c] == '#' {
	// 			grid[r][c] = 2
	// 		}
	// 	}
	// }

	r, c := 0, 0

	for i := range len(lines[0]) {
		if lines[0][i] == '.' {
			c = i
			break
		}
	}

	grid := make([][]int, rows)
	for i := range rows {
		grid[i] = make([]int, cols)
	}

	instructions := lines[len(lines)-1]
	num := 0
	direction := 0
	for _, chr := range instructions {
		if chr >= '0' && chr <= '9' {
			num *= 10
			num += int(chr - '0')
		} else if chr == 'R' {
			r, c = move(r, c, lines, grid, direction, num)
			direction = (direction + 1) % 4
			num = 0
		} else if chr == 'L' {
			r, c = move(r, c, lines, grid, direction, num)
			direction = (len(DIRECTIONS) + direction - 1) % 4
			num = 0
		} else {
			panic("invalid instruction")
		}
	}
	r, c = move(r, c, lines, grid, direction, num)

	//1 indexed
	r, c = r+1, c+1

	fmt.Println("93210 too low")
	return fmt.Sprint((1000 * r) + (4 * c) + direction)
}

func move(r, c int, lines []string, grid [][]int, direction, count int) (int, int) {
	for range count {
		r2 := r + DIRECTIONS[direction][0]
		c2 := c + DIRECTIONS[direction][1]
		for {
			// if r2 < 0 && direction == 3 {
			// 	r2 = len(grid) - 1
			// 	for c2 >= len(lines[r2]) || lines[r2][c2] == ' ' {
			// 		r2 -= 1
			// 	}
			// } else if c2 < 0 && direction == 2 {
			// 	c2 = len(lines[r2]) - 1
			// 	for lines[r2][c2] == ' ' {
			// 		c2 -= 1
			// 	}
			// } else if r2 >= len(grid) && direction == 1 {
			// 	r2 = 0
			// 	for c2 >= len(lines[r2]) || lines[r2][c2] == ' ' {
			// 		r2 += 1
			// 	}
			// } else if c2 >= len(lines[r2]) && direction == 0 {
			// 	// if direction != 0 {
			// 	// 	panic("WTF")
			// 	// }
			// 	c2 = 0
			// 	for lines[r2][c2] == ' ' {
			// 		c2 += 1
			// 	}
			// }

			if r2 >= 0 && c2 >= 0 && r2 < len(grid) && c2 < len(lines[r2]) && lines[r2][c2] != ' ' {
				break
			}

			r2 += DIRECTIONS[direction][0]
			c2 += DIRECTIONS[direction][1]

			r2 = (r2 + len(lines)) % len(lines)
			c2 = (c2 + len(lines[0])) % len(lines[0])
		}

		if lines[r2][c2] == '#' {
			break
		}
		if lines[r2][c2] != '.' {
			panic("fuk")
		}

		r, c = r2, c2
	}

	if lines[r][c] != '.' {
		panic("fuk")
	}

	return r, c
}

func part2(input string) string {
	return ""
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "22")
	if err != nil {
		log.Fatal(err)
	}

	// 	input = `        ...#
	//         .#..
	//         #...
	//         ....
	// ...#.......#
	// ........#...
	// ..#....#....
	// ..........#.
	//         ...#....
	//         .....#..
	//         .#......
	//         ......#.

	// 10R5L5R10L4R5L5`

	fmt.Println(part1(input))

	fmt.Println(part2(input))

	fmt.Println((time.Since(start)))
}
