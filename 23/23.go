package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"strings"
	"time"
)

type elf struct {
	i, j   int
	i2, j2 int
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	elves := []*elf{}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				elves = append(elves, &elf{
					i, j, 0, 0,
				})
			}
		}
	}

	dpWidth := 99999999
	dpOffset := 9999
	occupied := map[int]bool{}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	directionsToCheck := [4][3][2]int{
		{{-1, 0}, {-1, -1}, {-1, 1}},
		{{1, 0}, {1, -1}, {1, 1}},
		{{0, -1}, {1, -1}, {-1, -1}},
		{{0, 1}, {-1, 1}, {1, 1}}}

	getDpIdx := func(i, j int) int {
		return (i + dpOffset) + (j+dpOffset)*dpWidth
	}

	minI, minJ, maxI, maxJ := elves[0].i, elves[0].j, elves[0].i, elves[0].j
	for _, elf := range elves {
		minI, maxI, minJ, maxJ = min(minI, elf.i), max(maxI, elf.i), min(minJ, elf.j), max(maxJ, elf.j)
		dpIdx := getDpIdx(elf.i, elf.j)
		occupied[dpIdx] = true
	}

	// best := ((1 + maxI - minI) * (1 + maxJ - minJ)) - len(elves)
	ans := -1

	directionIdx := 0

	for range 10 {
		seens := map[int]*elf{}

		for _, elf := range elves {
			adj := false
			for i := -1; !adj && i < 2; i++ {
				for j := -1; j < 2; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if occupied[getDpIdx(elf.i+i, elf.j+j)] {
						adj = true
						break
					}
				}
			}
			if !adj {
				continue
			}

			for i := range len(directions) {
				curDirIdx := (directionIdx + i) % len(directions)
				wrong := false
				for _, toCheck := range directionsToCheck[curDirIdx] {
					if occupied[getDpIdx(elf.i+toCheck[0], elf.j+toCheck[1])] {
						wrong = true
						break
					}
				}
				if wrong {
					continue
				}
				elf.i2, elf.j2 = elf.i+directions[curDirIdx][0], elf.j+directions[curDirIdx][1]
				break
			}

			dpIdx := getDpIdx(elf.i2, elf.j2)
			if _, ok := seens[dpIdx]; ok {
				seens[dpIdx] = nil
				continue
			}
			seens[dpIdx] = elf
		}

		for _, elf := range seens {
			if elf == nil {
				continue
			}

			dpIdx1 := getDpIdx(elf.i, elf.j)
			dpIdx2 := getDpIdx(elf.i2, elf.j2)

			occupied[dpIdx1] = false
			occupied[dpIdx2] = true

			elf.i, elf.j = elf.i2, elf.j2
		}

		minI, minJ, maxI, maxJ := elves[0].i, elves[0].j, elves[0].i, elves[0].j
		for _, elf := range elves {
			minI, maxI, minJ, maxJ = min(minI, elf.i), max(maxI, elf.i), min(minJ, elf.j), max(maxJ, elf.j)
		}
		ans = ((1 + maxI - minI) * (1 + maxJ - minJ)) - len(elves)

		directionIdx += 1
		directionIdx %= len(directions)
	}
	return fmt.Sprint(ans)
}
func part2(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	elves := []*elf{}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				elves = append(elves, &elf{
					i, j, 0, 0,
				})
			}
		}
	}

	dpWidth := 999 * 999 * 999
	dpOffset := 999
	occupied := map[int]bool{}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	directionsToCheck := [4][3][2]int{
		{{-1, 0}, {-1, -1}, {-1, 1}},
		{{1, 0}, {1, -1}, {1, 1}},
		{{0, -1}, {1, -1}, {-1, -1}},
		{{0, 1}, {-1, 1}, {1, 1}}}

	getDpIdx := func(i, j int) int {
		return (i + dpOffset) + (j+dpOffset)*dpWidth
	}

	minI, minJ, maxI, maxJ := elves[0].i, elves[0].j, elves[0].i, elves[0].j
	for _, elf := range elves {
		minI, maxI, minJ, maxJ = min(minI, elf.i), max(maxI, elf.i), min(minJ, elf.j), max(maxJ, elf.j)
		dpIdx := getDpIdx(elf.i, elf.j)
		occupied[dpIdx] = true
	}

	// best := ((1 + maxI - minI) * (1 + maxJ - minJ)) - len(elves)
	ans := -1

	directionIdx := 0

	for round := range 10000000 {
		seens := map[int]*elf{}

		for _, elf := range elves {
			adj := false
			for i := -1; !adj && i < 2; i++ {
				for j := -1; j < 2; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if occupied[getDpIdx(elf.i+i, elf.j+j)] {
						adj = true
						break
					}
				}
			}
			if !adj {
				continue
			}

			found := false
			for i := range len(directions) {
				curDirIdx := (directionIdx + i) % len(directions)
				wrong := false
				for _, toCheck := range directionsToCheck[curDirIdx] {
					if occupied[getDpIdx(elf.i+toCheck[0], elf.j+toCheck[1])] {
						wrong = true
						break
					}
				}
				if wrong {
					continue
				}
				elf.i2, elf.j2 = elf.i+directions[curDirIdx][0], elf.j+directions[curDirIdx][1]
				found = true

				break
			}

			if !found {
				continue
			}

			dpIdx := getDpIdx(elf.i2, elf.j2)
			if _, ok := seens[dpIdx]; ok {
				seens[dpIdx] = nil
				continue
			}

			if elf.i2 == elf.i && elf.j2 == elf.j {
				panic("WTF")
			}
			seens[dpIdx] = elf
		}

		moved := false
		for _, elf := range seens {
			if elf == nil {
				continue
			}

			dpIdx1 := getDpIdx(elf.i, elf.j)
			dpIdx2 := getDpIdx(elf.i2, elf.j2)

			moved = true

			if elf.i2 == elf.i && elf.j2 == elf.j {
				panic("WTF")
			}

			occupied[dpIdx1] = false
			occupied[dpIdx2] = true

			elf.i, elf.j = elf.i2, elf.j2
		}
		if !moved {
			ans = round + 1
			break
		}

		minI, minJ, maxI, maxJ := elves[0].i, elves[0].j, elves[0].i, elves[0].j
		for _, elf := range elves {
			minI, maxI, minJ, maxJ = min(minI, elf.i), max(maxI, elf.i), min(minJ, elf.j), max(maxJ, elf.j)
		}

		directionIdx += 1
		directionIdx %= len(directions)
	}
	return fmt.Sprint(ans)
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "23")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))

	fmt.Println(part2(input))

	log.Println((time.Since(start)))
}
