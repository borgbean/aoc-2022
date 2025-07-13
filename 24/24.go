package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	empty byte = iota
	wall
	north = 2
	east  = 4
	south = 8
	west  = 16
)

var directions = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
	{0, 0},
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rows, cols := len(lines), len(lines[0])

	startCol, endCol := 0, 0

	grid := make([][]byte, rows)
	grid2 := make([][]byte, rows)
	dp := make([][]int, rows)
	for r, line := range lines {
		grid[r] = make([]byte, cols)
		grid2[r] = make([]byte, cols)
		dp[r] = make([]int, cols)
		for c, val := range line {
			switch val {
			case '#':
				grid[r][c] = wall
			case '.':
				grid[r][c] = empty
				if r == 0 {
					startCol = c
				} else if r == rows-1 {
					endCol = c
				}
			case '>':
				grid[r][c] = east
			case '<':
				grid[r][c] = west
			case '^':
				grid[r][c] = north
			case 'v':
				grid[r][c] = south
			}
		}
	}

	move1 := func(dir, r, c int) {
		for {
			r = (r + directions[dir][0] + rows) % rows
			c = (c + directions[dir][1] + cols) % cols
			if grid2[r][c] != wall {
				grid2[r][c] |= 2 << dir
				return
			}
		}
	}

	advance := func() {
		for r, line := range grid {
			for c := range line {
				if grid[r][c] == wall {
					grid2[r][c] = wall
				} else {
					grid2[r][c] = empty
				}
			}
		}

		for r, line := range grid {
			for c, val := range line {
				if val == empty || val == wall {
					continue
				}

				for ord := range 4 {
					if grid[r][c]&(2<<ord) > 0 {
						move1(ord, r, c)
					}
				}
			}
		}

		grid, grid2 = grid2, grid
	}

	q1 := [][2]int{{0, startCol}}
	q2 := [][2]int{}
	steps := 0
	advance()

	for len(q1) > 0 {
		cur := q1[len(q1)-1]
		q1 = q1[:len(q1)-1]

		for _, dir := range directions {
			r2, c2 := cur[0]+dir[0], cur[1]+dir[1]

			if r2 < 0 || c2 < 0 || r2 >= rows || c2 >= cols || grid[r2][c2] != empty {
				continue
			}

			if r2 == rows-1 && c2 == endCol {
				return fmt.Sprint(steps + 1)
			}

			itVal := int(empty) + steps + 1
			if dp[r2][c2] >= itVal {
				continue
			}

			dp[r2][c2] = itVal
			q2 = append(q2, [2]int{r2, c2})
		}

		if len(q1) < 1 {
			q1, q2 = q2, q1
			advance()
			steps += 1
		}
	}

	return "broken"
}

func part2(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rows, cols := len(lines), len(lines[0])

	startCol, endCol := 0, 0

	grid := make([][]byte, rows)
	blizzards := [][3]int{}
	dp := [3][][]int{
		make([][]int, rows),
		make([][]int, rows),
		make([][]int, rows),
	}
	for r, line := range lines {
		grid[r] = make([]byte, cols)
		for _, arr := range dp {
			arr[r] = make([]int, cols)
		}
		for c, val := range line {
			switch val {
			case '#':
				grid[r][c] = wall
			case '.':
				grid[r][c] = empty
				if r == 0 {
					startCol = c
				} else if r == rows-1 {
					endCol = c
				}
			case '>':
				grid[r][c] = east
				blizzards = append(blizzards, [3]int{r, c, 1})
			case '<':
				grid[r][c] = west
				blizzards = append(blizzards, [3]int{r, c, 3})
			case '^':
				grid[r][c] = north
				blizzards = append(blizzards, [3]int{r, c, 0})
			case 'v':
				grid[r][c] = south
				blizzards = append(blizzards, [3]int{r, c, 2})
			}
		}
	}

	move1 := func(dir, r, c int) (int, int) {
		for {
			r = (r + directions[dir][0] + rows) % rows
			c = (c + directions[dir][1] + cols) % cols
			if grid[r][c] != wall {
				return r, c
			}
		}
	}

	advance := func() {
		for _, blizzard := range blizzards {
			r, c := blizzard[0], blizzard[1]
			grid[r][c] = 0
		}
		for idx := range blizzards {
			r, c, direction := blizzards[idx][0], blizzards[idx][1], blizzards[idx][2]
			r, c = move1(direction, r, c)
			blizzards[idx][0] = r
			blizzards[idx][1] = c
			grid[r][c] |= 2 << direction
		}
	}

	q1 := [][3]int{{0, startCol, 0}}
	q2 := [][3]int{}
	steps := 0
	advance()

	for len(q1) > 0 {
		cur := q1[len(q1)-1]
		r1, c1, step := cur[0], cur[1], cur[2]
		q1 = q1[:len(q1)-1]

		for _, dir := range directions {
			r2, c2, step2 := r1+dir[0], c1+dir[1], step

			if r2 < 0 || c2 < 0 || r2 >= rows || c2 >= cols || grid[r2][c2] != empty {
				continue
			}

			if r2 == 0 && c2 == startCol {
				if step2 == 1 {
					step2 += 1
				}
			}
			if r2 == rows-1 && c2 == endCol {
				if step2 == 0 {
					step2 += 1
				} else if step2 == 2 {
					return fmt.Sprint(steps + 1)
				}
			}

			itVal := int(empty) + steps + 1
			if dp[step2][r2][c2] >= itVal {
				continue
			}

			dp[step2][r2][c2] = itVal
			q2 = append(q2, [3]int{r2, c2, step2})
		}

		if len(q1) < 1 {
			q1, q2 = q2, q1
			advance()
			steps += 1
		}
	}

	return "broken"
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(input))

	fmt.Println(part2(input))

	log.Println((time.Since(start)))
}
