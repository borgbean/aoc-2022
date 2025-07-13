package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"time"
)

func part1(input string, row int) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	instructionRex := regexp.MustCompile(`.*x=(-?\d+), y=(-?\d+).*x=(-?\d+), y=(-?\d+).*`)

	blockages := [][2]int{}
	beacons := []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		match := instructionRex.FindStringSubmatch(line)

		x0S, y0S, x1S, y1S := match[1], match[2], match[3], match[4]

		x0, err := strconv.Atoi(x0S)
		if err != nil {
			log.Panic("bad line", line, err)
		}

		x1, err := strconv.Atoi(x1S)
		if err != nil {
			log.Panic("bad line", line, err)
		}
		y0, err := strconv.Atoi(y0S)
		if err != nil {
			log.Panic("bad line", line, err)
		}
		y1, err := strconv.Atoi(y1S)
		if err != nil {
			log.Panic("bad line", line, err)
		}

		dist := max(x0-x1, x1-x0) + max(y0-y1, y1-y0)

		width := dist*2 + 1
		verticalDist := max(y0-row, row-y0)

		blocked := max(0, width-(verticalDist*2))
		if blocked > 0 {
			blockages = append(blockages, [2]int{x0 - blocked/2, x0 + blocked/2})
		}

		if y1 == row {
			beacons = append(beacons, x1)
		}
	}

	slices.SortFunc(blockages, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	result := 0

	merged := [][2]int{blockages[0]}

	for _, blockage := range blockages[1:] {
		if blockage[0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], blockage[1])
		} else {
			merged = append(merged, blockage)
		}
	}

	for _, interval := range merged {
		result += 1 + interval[1] - interval[0]
	}

	slices.Sort(beacons)

	intervalIdx := 0
	for i, beacon := range beacons {
		if i > 0 && beacon == beacons[i-1] {
			continue
		}

		for intervalIdx < len(merged) && merged[intervalIdx][1] < beacon {
			intervalIdx++
		}
		if intervalIdx > len(merged) {
			break
		}

		if merged[intervalIdx][0] <= beacon && merged[intervalIdx][1] >= beacon {
			result -= 1
		}
	}

	return fmt.Sprint(result)
}

type sensor struct {
	x, y   int
	radius int
}

func part2(input string, bound int) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	instructionRex := regexp.MustCompile(`.*x=(-?\d+), y=(-?\d+).*x=(-?\d+), y=(-?\d+).*`)

	sensors := []sensor{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		match := instructionRex.FindStringSubmatch(line)

		x0S, y0S, x1S, y1S := match[1], match[2], match[3], match[4]

		x0, err := strconv.Atoi(x0S)
		if err != nil {
			log.Panic("bad line", line, err)
		}

		x1, err := strconv.Atoi(x1S)
		if err != nil {
			log.Panic("bad line", line, err)
		}
		y0, err := strconv.Atoi(y0S)
		if err != nil {
			log.Panic("bad line", line, err)
		}
		y1, err := strconv.Atoi(y1S)
		if err != nil {
			log.Panic("bad line", line, err)
		}

		dist := max(x0-x1, x1-x0) + max(y0-y1, y1-y0)

		sensors = append(sensors, sensor{x0, y0, dist})
	}

	for i := range sensors {
		//go around the perimeter

		/**
			 0
			1*1
		   2*_*2
			3*3
		     4
		*/

		toCompare := []sensor{}
		for j := range sensors {
			if i == j {
				continue
			}

			xDiff := sensors[j].x - sensors[i].x
			yDiff := sensors[j].y - sensors[i].y
			diff := max(-xDiff, xDiff) + max(-yDiff, yDiff)

			if diff-(1+sensors[i].radius+sensors[j].radius) <= 1 {
				toCompare = append(toCompare, sensors[j])
			}
		}

		check := func(y, x int) bool {
			if x <= 0 || y <= 0 || x >= bound || y >= bound {
				return false
			}
			for _, s2 := range toCompare {
				// for _, s2 := range sensors {
				dist := max(s2.x-x, x-s2.x) + max(s2.y-y, y-s2.y)
				if dist <= s2.radius {
					return false
				}
			}

			return true
		}

		sensor := sensors[i]

		y := (sensor.y - sensor.radius) - 1

		radius := 0
		if check(y, sensor.x) {
			return fmt.Sprintf("%v:%v", sensor.x, y)
		}
		y += 1

		for y <= sensor.y {
			radius += 1

			if check(y, sensor.x-radius) {
				return fmt.Sprintf("%v:%v", sensor.x-radius, y)
			}
			if check(y, sensor.x+radius) {
				return fmt.Sprintf("%v:%v", sensor.x+radius, y)
			}
			y += 1
		}

		yEnd := sensor.y + sensor.radius + 1
		for y <= bound && y < yEnd {
			radius -= 1

			if check(y, sensor.x-radius) {
				return fmt.Sprintf("%v:%v", sensor.x-radius, y)
			}
			if check(y, sensor.x+radius) {
				return fmt.Sprintf("%v:%v", sensor.x+radius, y)
			}
			y += 1
		}
		if check(y, sensor.x) {
			return fmt.Sprintf("%v:%v", sensor.x, y)
		}
	}

	return "wrong!"
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "15")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input, 2000000))
	fmt.Println(part2(input, 4000000))

	log.Println((time.Since(start)))
}
