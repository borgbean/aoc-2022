package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type position struct{ i, j int }
type movement struct{ i, j int }

var DIRECTIONS = map[byte]movement{
	'U': {-1, 0},
	'D': {1, 0},
	'L': {0, -1},
	'R': {0, 1},
}

func part1(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	i1, j1 := 0, 0
	i2, j2 := 0, 0

	seen := map[uint64]struct{}{}

	seen[0] = struct{}{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			log.Panic("bad input", line, err)
		}

		step := DIRECTIONS[direction]
		iStep, jStep := step.i, step.j

		for range distance {
			i1 += iStep
			j1 += jStep

			if max(i1-i2, i2-i1) < 2 && max(j1-j2, j2-j1) < 2 {
				continue
			}

			if i1 != i2 && j1 != j2 {
				//move diagonally
				if i2 < i1 {
					i2 += 1
				} else {
					i2 -= 1
				}

				if j2 < j1 {
					j2 += 1
				} else {
					j2 -= 1
				}
			} else {
				i2 += iStep
				j2 += jStep
			}

			dpIdx := uint64(uint32(i2)) << 32
			dpIdx = uint64(uint32(j2)) | dpIdx

			seen[dpIdx] = struct{}{}
		}
	}

	return fmt.Sprint(len(seen))

}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	positions := [10]*position{}
	for i := 0; i < len(positions); i++ {
		positions[i] = &position{}
	}

	seen := map[uint64]struct{}{}

	seen[0] = struct{}{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			log.Panic("bad input", line, err)
		}

		step := DIRECTIONS[direction]
		iStep, jStep := step.i, step.j

		for range distance {
			curIStep, curJStep := iStep, jStep

			positions[0].i += iStep
			positions[0].j += jStep
			prevPos := positions[0]

			for _, pos := range positions[1:] {
				if max(prevPos.i-pos.i, pos.i-prevPos.i) < 2 && max(prevPos.j-pos.j, pos.j-prevPos.j) < 2 {
					break
				}

				if prevPos.i != pos.i && prevPos.j != pos.j {
					//move diagonally
					if pos.i < prevPos.i {
						curIStep = 1
					} else {
						curIStep = -1
					}

					if pos.j < prevPos.j {
						curJStep = 1
					} else {
						curJStep = -1
					}
				} else {
					if pos.i == prevPos.i {
						curIStep = 0
					} else {
						curJStep = 0
					}
				}
				pos.i += curIStep
				pos.j += curJStep

				prevPos = pos
			}

			lastPos := positions[len(positions)-1]
			dpIdx := uint64(uint32(lastPos.i)) << 32
			dpIdx = uint64(uint32(lastPos.j)) | dpIdx

			seen[dpIdx] = struct{}{}
		}
	}

	return fmt.Sprint(len(seen))
}

func main() {
	input, err := aoc.GetInput("2022", "9")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(part1(input))
	log.Println(part2(input))
}
