package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var SIDE_OFFSETS = [][3]int{
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	seenSides := map[int]bool{}
	for _, line := range lines {
		split := strings.Split(line, ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("bad line", line, err)
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("bad line", line, err)
		}
		z, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal("bad line", line, err)
		}

		for _, offset := range SIDE_OFFSETS {
			x2, y2, z2 := (2*x)+offset[0], (2*y)+offset[1], (2*z)+offset[2]

			dpIdx := x2*1000 + y2*1000*1000 + z2

			seenSides[dpIdx] = true
		}
	}

	totalSides := len(lines) * 6
	distinctSides := len(seenSides)
	result := distinctSides - (totalSides - distinctSides)

	return fmt.Sprint("sides: ", result)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	cubes := map[int]bool{}
	for _, line := range lines {
		split := strings.Split(line, ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("bad line", line, err)
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("bad line", line, err)
		}
		z, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal("bad line", line, err)
		}

		x += 1
		y += 1
		z += 1
		{
			dpIdx := x*1000*1000 + y*1000 + z
			cubes[dpIdx] = true
		}
	}

	grid := [45][45][45]int{}
	//0 <- null, 1 <- inside, 2 <- outside
	groupParents := []int{0, 1, 2}

	bfs := func(x, y, z int) int {
		s := [][3]int{{x, y, z}}
		s2 := [][3]int{}

		group := len(groupParents)

		for len(s) > 0 {
			cur := s[len(s)-1]
			s = s[:len(s)-1]

			x, y, z := cur[0], cur[1], cur[2]

			for _, offset := range SIDE_OFFSETS {
				x2, y2, z2 := x+offset[0], y+offset[1], z+offset[2]

				//out of bounds?
				if x2 < 0 || y2 < 0 || z2 < 0 || x2 > 44 || y2 > 44 || z2 > 44 {
					//outside
					groupParents = append(groupParents, 2)
					// group = 2
					group = groupParents[group]

					return group
				}

				if grid[x2][y2][z2] == group {
					continue
				}
				dpIdx := x2*1000*1000 + y2*1000 + z2
				if cubes[dpIdx] {
					//can't go INSIDE cube
					continue
				}

				if grid[x2][y2][z2] != 0 {
					//we found the answer
					groupParents = append(groupParents, groupParents[grid[x2][y2][z2]])
					group = groupParents[group]
					return group
				}

				grid[x2][y2][z2] = group
				s2 = append(s2, [3]int{x2, y2, z2})
			}

			if len(s) < 1 {
				s, s2 = s2, s
			}
		}

		//inside
		groupParents = append(groupParents, 1)
		return 1
	}

	result := 0

	for _, line := range lines {
		split := strings.Split(line, ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("bad line", line, err)
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("bad line", line, err)
		}
		z, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal("bad line", line, err)
		}

		x += 1
		y += 1
		z += 1

		for _, offset := range SIDE_OFFSETS {
			x2, y2, z2 := x+offset[0], y+offset[1], z+offset[2]

			if x2 < 0 || y2 < 0 || z2 < 0 {
				continue
			}

			dpIdx := x2*1000*1000 + y2*1000 + z2
			if cubes[dpIdx] {
				continue
			}

			if bfs(x2, y2, z2) == 2 {
				result += 1
			}
		}

	}

	return fmt.Sprint("sides: ", result)
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "18")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(part1(input))
	log.Println(part2(input))

	log.Println((time.Since(start)))
}
