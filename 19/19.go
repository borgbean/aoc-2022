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

type state struct {
	ore, clay, obsidian, geode, t, oreRobot, clayRobot, obsidianRobot, geodeRobot int
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rex := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

	totalTime := 24
	ret := 0

	for _, line := range lines {
		best := 0
		match := rex.FindStringSubmatch(line)
		blueprint := parseInt(match[1])
		oreRobotCost := parseInt(match[2])
		clayRobotCost := parseInt(match[3])
		obsRobotCostOre := parseInt(match[4])
		obsRobotCostClay := parseInt(match[5])
		geodeRobotCostOre := parseInt(match[6])
		geodeRobotCostObs := parseInt(match[7])

		s := []state{}

		s = append(s, state{0, 0, 0, 0, 0, 1, 0, 0, 0})

		maxOreRobots := max(clayRobotCost, obsRobotCostOre, geodeRobotCostOre)
		maxObsRobots := geodeRobotCostObs
		maxClayRobots := obsRobotCostClay

		for len(s) > 0 {
			cur := s[len(s)-1]
			s = s[:len(s)-1]

			if cur.obsidianRobot > 0 && best > 0 {
				curNext := cur
				for t := cur.t; t < totalTime; t++ {
					curNext = jumpTo(curNext, curNext, curNext.t, curNext.t+1)
					curNext.geodeRobot++
				}
				if curNext.geode <= best {
					continue
				}
			}

			if cur.t > totalTime {
				continue
			}

			best = max(best, cur.geode, -blueprint)

			if cur.t == totalTime {
				continue
			}

			//buy an ore robot if possible
			if cur.oreRobot < maxOreRobots {
				duration := max(0, ((cur.oreRobot-1)+oreRobotCost-cur.ore)/cur.oreRobot)
				curNext := jumpTo(cur, cur, cur.t, 1+duration+cur.t)
				curNext.ore -= oreRobotCost
				curNext.oreRobot += 1
				s = append(s, curNext)
			}

			//buy a clay robot if possible
			if cur.clayRobot < maxClayRobots {
				duration := max(0, ((cur.oreRobot-1)+clayRobotCost-cur.ore)/cur.oreRobot)
				curNext := jumpTo(cur, cur, cur.t, 1+cur.t+duration)
				curNext.ore -= clayRobotCost
				curNext.clayRobot += 1
				s = append(s, curNext)
			}

			//buy a obsidian robot if possible
			if cur.clayRobot > 0 && cur.obsidianRobot < maxObsRobots {
				duration := max(0,
					((cur.oreRobot-1)+obsRobotCostOre-cur.ore)/cur.oreRobot,
					((cur.clayRobot-1)+obsRobotCostClay-cur.clay)/cur.clayRobot)

				curNext := jumpTo(cur, cur, cur.t, cur.t+1+duration)

				curNext.ore -= obsRobotCostOre
				curNext.clay -= obsRobotCostClay
				curNext.obsidianRobot += 1

				s = append(s, curNext)
			}

			//buy a geode robot if possible
			if cur.obsidianRobot > 0 {
				duration := max(0,
					((cur.oreRobot-1)+geodeRobotCostOre-cur.ore)/cur.oreRobot,
					((cur.obsidianRobot-1)+geodeRobotCostObs-cur.obsidian)/cur.obsidianRobot)

				curNext := jumpTo(cur, cur, cur.t, cur.t+1+duration)

				curNext.ore -= geodeRobotCostOre
				curNext.obsidian -= geodeRobotCostObs
				curNext.geodeRobot += 1

				s = append(s, curNext)
			}
		}

		ret += best * blueprint

	}

	return fmt.Sprint(ret)
}

func jumpTo(s1, s2 state, t0, t1 int) state {
	s2.ore += s1.oreRobot * (t1 - t0)
	s2.clay += s1.clayRobot * (t1 - t0)
	s2.obsidian += s1.obsidianRobot * (t1 - t0)
	s2.geode += s1.geodeRobot * (t1 - t0)

	s2.t = t1

	return s2
}

func part2(input string) string {
	lines := strings.Split(input, "\n")

	lines = lines[:3]

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rex := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

	totalTime := 32
	ret := 1

	for _, line := range lines {
		best := 0
		match := rex.FindStringSubmatch(line)
		// blueprint := parseInt(match[1])
		oreRobotCost := parseInt(match[2])
		clayRobotCost := parseInt(match[3])
		obsRobotCostOre := parseInt(match[4])
		obsRobotCostClay := parseInt(match[5])
		geodeRobotCostOre := parseInt(match[6])
		geodeRobotCostObs := parseInt(match[7])

		s := []state{}

		s = append(s, state{0, 0, 0, 0, 0, 1, 0, 0, 0})

		maxOreRobots := max(clayRobotCost, obsRobotCostOre, geodeRobotCostOre)
		maxObsRobots := geodeRobotCostObs
		maxClayRobots := obsRobotCostClay

		for len(s) > 0 {
			cur := s[len(s)-1]
			s = s[:len(s)-1]

			if cur.obsidianRobot > 0 && best > 0 {
				curNext := cur
				for t := cur.t; t < totalTime; t++ {
					curNext = jumpTo(curNext, curNext, curNext.t, curNext.t+1)
					curNext.geodeRobot++
				}
				if curNext.geode <= best {
					continue
				}
			}

			if cur.t > totalTime {
				continue
			}

			best = max(best, cur.geode)

			if cur.t == totalTime {
				continue
			}

			//buy an ore robot if possible
			if cur.oreRobot < maxOreRobots {
				duration := max(0, ((cur.oreRobot-1)+oreRobotCost-cur.ore)/cur.oreRobot)
				curNext := jumpTo(cur, cur, cur.t, 1+duration+cur.t)
				curNext.ore -= oreRobotCost
				curNext.oreRobot += 1

				if curNext.t <= totalTime {
					s = append(s, curNext)
				}
			}

			//buy a clay robot if possible
			if cur.clayRobot < maxClayRobots {
				duration := max(0, ((cur.oreRobot-1)+clayRobotCost-cur.ore)/cur.oreRobot)
				curNext := jumpTo(cur, cur, cur.t, 1+cur.t+duration)
				curNext.ore -= clayRobotCost
				curNext.clayRobot += 1

				if curNext.t <= totalTime {
					s = append(s, curNext)
				}
			}

			//buy a obsidian robot if possible
			if cur.clayRobot > 0 &&
				cur.obsidianRobot < maxObsRobots {
				duration := max(0,
					((cur.oreRobot-1)+obsRobotCostOre-cur.ore)/cur.oreRobot,
					((cur.clayRobot-1)+obsRobotCostClay-cur.clay)/cur.clayRobot)

				curNext := jumpTo(cur, cur, cur.t, cur.t+1+duration)

				curNext.ore -= obsRobotCostOre
				curNext.clay -= obsRobotCostClay
				curNext.obsidianRobot += 1

				if curNext.t <= totalTime {
					s = append(s, curNext)
				}
			}

			// buy a geode robot if possible
			if cur.obsidianRobot > 0 && cur.ore >= geodeRobotCostOre {
				duration := max(0,
					((cur.oreRobot-1)+geodeRobotCostOre-cur.ore)/cur.oreRobot,
					((cur.obsidianRobot-1)+geodeRobotCostObs-cur.obsidian)/cur.obsidianRobot)

				curNext := jumpTo(cur, cur, cur.t, cur.t+1+duration)

				curNext.ore -= geodeRobotCostOre
				curNext.obsidian -= geodeRobotCostObs
				curNext.geodeRobot += 1

				if curNext.t <= totalTime {
					s = append(s, curNext)
				}
			}
		}

		ret *= best

	}

	return fmt.Sprint(ret)
}

func parseInt(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		log.Panic("Failed to parse int", s, err)
	}

	return ret
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "19")
	if err != nil {
		log.Fatal(err)
	}

	// 	input = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
	// Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
	// `

	log.Println(part1(input))

	log.Println(part2(input))

	log.Println((time.Since(start)))
}

/*

(

	(
		(, )?

		([^,]+)
	)+

)


*/
