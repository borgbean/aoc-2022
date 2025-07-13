package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"
)

func part1(input string) string {
	const shapesRaw = `####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##`

	shapesStrings := strings.Split(shapesRaw, "\n\n")

	shapes := [][]string{}

	for _, shapeS := range shapesStrings {
		shape := strings.Split(shapeS, "\n")
		slices.Reverse(shape)
		shapes = append(shapes, shape)
	}

	field := [][7]bool{}
	top := 0
	shapeIdx := 0
	moveIdx := 0

	for input[len(input)-1] < '<' || input[len(input)-1] > '>' {
		input = input[:len(input)-1]
	}

	for range 2022 {
		shape := shapes[shapeIdx]
		shapeIdx = (shapeIdx + 1) % len(shapes)
		start := top + 3

		for (start + len(shape)) > len(field) {
			field = append(field, [7]bool{})
		}

		var newTop int
		newTop, moveIdx = place(shape, start, field, input, moveIdx)

		top = max(top, newTop)
	}

	return fmt.Sprint(top)
}

func part2(input string) string {
	const shapesRaw = `####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##`

	shapesStrings := strings.Split(shapesRaw, "\n\n")

	shapes := [][]string{}

	for _, shapeS := range shapesStrings {
		shape := strings.Split(shapeS, "\n")
		slices.Reverse(shape)
		shapes = append(shapes, shape)
	}

	field := [][7]bool{}
	top := 0
	shapeIdx := 0
	moveIdx := 0

	for input[len(input)-1] < '<' || input[len(input)-1] > '>' {
		input = input[:len(input)-1]
	}

	dp := map[string]int{}
	dpToHeight := map[int]int{}

	var dp1, dp2 int

	for it := range 5000 {
		it = it

		shape := shapes[shapeIdx]
		shapeIdx = (shapeIdx + 1) % len(shapes)
		start := top + 3

		for (start + len(shape)) > len(field) {
			field = append(field, [7]bool{})
		}

		var newTop int
		newTop, moveIdx = place(shape, start, field, input, moveIdx)

		top = max(top, newTop)

		dpIdx := fmt.Sprint(shapeIdx, "-", moveIdx, "-")
		for i := range 7 {
			for j := top; j >= 0; j-- {
				if field[j][i] {
					dpIdx += fmt.Sprint(top-j, "-")
					break
				}
			}
		}

		dpToHeight[it] = top
		if dp[dpIdx] > 0 {
			dp1, dp2 = dp[dpIdx], it
			break
		}
		dp[dpIdx] = it
	}

	target := 1000000000000

	cycleLen := dp2 - dp1
	cycles := target / cycleLen
	result := (dpToHeight[dp2] - dpToHeight[dp1]) * cycles
	target -= cycleLen * cycles
	result += dpToHeight[target-1]

	return fmt.Sprint(result)
}

func place(shape []string, top int, field [][7]bool, moves string, moveIdx int) (int, int) {
	left := 2
	pos := top

	for {
		if moves[moveIdx] == '<' {
			//left
			if !intersects(shape, pos, left-1, field) {
				left -= 1
			}
		} else {
			//right
			if !intersects(shape, pos, left+1, field) {
				left += 1
			}
		}

		moveIdx = (moveIdx + 1) % len(moves)

		if intersects(shape, pos-1, left, field) {
			break
		}
		pos--
	}

	drawShape(shape, pos, left, field)

	top = pos + len(shape)
	return top, moveIdx
}

func intersects(shape []string, row, col int, field [][7]bool) bool {
	if row < 0 || col < 0 || (col+len(shape[0])) > len(field[row]) {
		return true
	}

	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] != '#' {
				continue
			}

			if field[row+i][j+col] {
				return true
			}
		}
	}

	return false
}

func drawShape(shape []string, row, col int, field [][7]bool) {
	for i := range shape {
		for j := range len(shape[i]) {
			if shape[i][j] == '#' {
				field[row+i][j+col] = true
			}
		}
	}
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "17")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))

	log.Println((time.Since(start)))
}
