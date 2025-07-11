package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"strings"
	"time"
)

var DIRECTIONS = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

const (
	east int = iota
	south
	west
	north
)

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rows := len(lines) - 2
	cols := len(lines[len(lines)-3])

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
			r, c = move(r, c, lines, rows, cols, direction, num)
			direction = (direction + 1) % 4
			num = 0
		} else if chr == 'L' {
			r, c = move(r, c, lines, rows, cols, direction, num)
			direction = (len(DIRECTIONS) + direction - 1) % 4
			num = 0
		} else {
			panic("invalid instruction")
		}
	}
	r, c = move(r, c, lines, rows, cols, direction, num)

	//1 indexed
	r, c = r+1, c+1

	return fmt.Sprint((1000 * r) + (4 * c) + direction)
}

func move(r, c int, lines []string, rows, cols, direction, count int) (int, int) {
	for range count {
		r2 := r + DIRECTIONS[direction][0]
		c2 := c + DIRECTIONS[direction][1]
		for {

			if r2 >= 0 && c2 >= 0 && r2 < rows && c2 < len(lines[r2]) && lines[r2][c2] != ' ' {
				break
			}

			r2 += DIRECTIONS[direction][0]
			c2 += DIRECTIONS[direction][1]

			r2 = (r2 + rows) % rows
			c2 = (c2 + len(lines[0])) % len(lines[0])
		}

		if lines[r2][c2] == '#' {
			break
		}
		if lines[r2][c2] != '.' {
			panic("uhoh")
		}

		r, c = r2, c2
	}

	if lines[r][c] != '.' {
		panic("uhoh")
	}

	return r, c
}

func part2(input string) string {

	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rows := len(lines) - 2
	cols := len(lines[0])

	r, c := 0, 0

	for i := range len(lines[0]) {
		if lines[0][i] == '.' {
			c = i
			break
		}
	}

	lines2 := make([][]byte, rows)
	for i := range rows {
		lines2[i] = make([]byte, cols)
	}

	instructions := lines[len(lines)-1]
	num := 0
	direction := 0
	for _, chr := range instructions {
		if chr >= '0' && chr <= '9' {
			num *= 10
			num += int(chr - '0')
		} else if chr == 'R' {
			r, c, direction = move2(r, c, lines, rows, cols, direction, num)
			direction = (direction + 1) % 4
			num = 0
		} else if chr == 'L' {
			r, c, direction = move2(r, c, lines, rows, cols, direction, num)
			direction = (len(DIRECTIONS) + direction - 1) % 4
			num = 0
		} else {
			panic("invalid instruction")
		}
	}
	r, c, direction = move2(r, c, lines, rows, cols, direction, num)

	//1 indexed
	r, c = r+1, c+1
	return fmt.Sprint((1000 * r) + (4 * c) + direction)
}

func move2(r, c int, lines []string, rows, cols, direction, count int) (int, int, int) {
	for range count {
		r2 := r + DIRECTIONS[direction][0]
		c2 := c + DIRECTIONS[direction][1]
		direction2 := direction

		if !(r2 >= 0 && c2 >= 0 && r2 < rows && c2 < len(lines[r2]) && lines[r2][c2] != ' ') {
			//need to teleport to another side of the cube

			//invert action
			r2 -= DIRECTIONS[direction][0]
			c2 -= DIRECTIONS[direction][1]

			//we treat the input as a 4x3 grid, even though only 6 of the spots are actually used
			gridRow := r / (rows / 4)
			gridCol := c / (cols / 3)

			tileWidth, tileHeight := (cols/3)-1, ((rows / 4) - 1)

			tileRow := r % (tileHeight + 1)
			tileCol := c % (tileWidth + 1)

			var newTileRow, newTileCol, newTile int
			if gridRow == 0 && gridCol == 1 {
				//tile 1
				if direction == west {
					newTile = 6
					newTileRow = tileHeight - tileRow
					newTileCol = 0

					//dirction flipped left -> right
					direction2 = east
				} else if direction == north {
					newTile = 9
					newTileRow = tileCol
					newTileCol = 0

					//dirction flipped up -> right
					direction2 = east
				}
			} else if gridRow == 0 && gridCol == 2 {
				//tile 2
				if direction == north {
					newTile = 9
					newTileRow = tileHeight
					newTileCol = tileCol

					//dirction flipped up -> up
					direction2 = north
				} else if direction == east {
					newTile = 7
					newTileRow = tileHeight - tileRow
					newTileCol = tileWidth

					//dirction flipped right -> left
					direction2 = west
				} else if direction == south {
					newTile = 4
					newTileRow = tileCol
					newTileCol = tileWidth

					//direction flipped down -> left
					direction2 = west
				}
			} else if gridRow == 1 && gridCol == 1 {
				//tile 5
				if direction == west {
					newTile = 6
					newTileRow = 0
					newTileCol = tileRow

					//direction flipped left -> down
					direction2 = south
				} else if direction == east {
					newTile = 2
					newTileRow = tileHeight
					newTileCol = tileRow

					//direction flipped right -> up
					direction2 = north
				}
			} else if gridRow == 2 && gridCol == 0 {
				//tile 6
				if direction == west {
					newTile = 1
					newTileRow = tileHeight - tileRow
					newTileCol = 0

					//direction flipped left -> right
					direction2 = east
				} else if direction == north {
					newTile = 4
					newTileRow = tileCol
					newTileCol = 0

					//direction flipped up -> right
					direction2 = east
				}
			} else if gridRow == 2 && gridCol == 1 {
				//tile 7
				if direction == east {
					newTile = 2
					newTileRow = tileHeight - tileRow
					newTileCol = tileWidth

					//direction flipped right -> left
					direction2 = west
				} else if direction == south {
					newTile = 9
					newTileRow = tileCol
					newTileCol = tileWidth

					//direction flipped down -> left
					direction2 = west
				}
			} else if gridRow == 3 && gridCol == 0 {
				//tile 9
				if direction == west {
					newTile = 1
					newTileRow = 0
					newTileCol = tileRow

					//direction flipped left -> down
					direction2 = south
				} else if direction == south {
					newTile = 2
					newTileRow = 0
					newTileCol = tileCol

					//direction flipped down -> down
					direction2 = south
				} else if direction == east {
					newTile = 7
					newTileRow = tileHeight
					newTileCol = tileRow

					//direction flipped right -> up
					direction2 = 3
				}
			}

			if newTile == 0 {
				panic("unhandled direction")
			}
			r2 = newTileRow + ((newTile / 3) * (tileHeight + 1))
			c2 = newTileCol + ((newTile % 3) * (tileWidth + 1))

			if lines[r2][c2] != '.' && lines[r2][c2] != '#' {
				panic("uhoh")
			}
		}

		if lines[r2][c2] == '#' {
			break
		}
		if lines[r2][c2] != '.' {
			panic("uhoh")
		}

		r, c, direction = r2, c2, direction2
	}

	if lines[r][c] != '.' {
		panic("uhoh")
	}

	return r, c, direction
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "22")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))

	fmt.Println(part2(input))

	fmt.Println((time.Since(start)))
}
