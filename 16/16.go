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

type vertex struct {
	rate int
	adj  []string
}

type state struct {
	mask   int64
	v      int
	rate   int
	flowed int
	t      int
}

func part1(input string) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	rex := regexp.MustCompile(`Valve ([^\s]+) has flow rate=(\d+); tunnels? leads? to valves? ((?:[^,]+,?)+)`)

	adjList := []vertex{}
	vertexNameToId := map[string]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := rex.FindStringSubmatch(line)

		v0, rate, valves := matches[1], matches[2], strings.Split(matches[3], ", ")

		rateVal, err := strconv.Atoi(rate)
		if err != nil {
			log.Panic("Failed to read rate for line", line, err)
		}

		vertexNameToId[v0] = len(vertexNameToId)
		adjList = append(adjList, vertex{
			rateVal, valves,
		})
	}

	adjMatrix := make([][]int, len(adjList))

	for vStart := range len(adjList) {
		adjMatrix[vStart] = make([]int, len(adjList))

		//[mask, v]
		dist := 0
		q := [][2]int{{1 << vStart, vStart}}
		q2 := [][2]int{}

		for len(q) > 0 {
			cur := q[len(q)-1]
			q = q[:len(q)-1]
			mask, v := cur[0], cur[1]

			for _, v2s := range adjList[v].adj {
				v2 := vertexNameToId[v2s]

				if v2 == vStart || (adjMatrix[vStart][v2] > 0 && adjMatrix[vStart][v2] <= dist) {
					continue
				}

				adjMatrix[vStart][v2] = dist + 1
				mask2 := mask | (1 << v2)
				q2 = append(q2, [2]int{mask2, v2})
			}

			if len(q) < 1 {
				dist += 1
				q, q2 = q2, q
			}
		}
	}

	s := []state{}

	{
		vStart := vertexNameToId["AA"]
		s = append(s, state{
			1 << vStart,
			vStart, //TODO this only kinda works because v0 has rate 0
			0,
			0,
			0,
		})
	}

	limit := 30
	best := 0

	dp := map[int64]int{}

	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]

		best = max(best, cur.flowed+(limit-cur.t)*cur.rate)

		for v2 := range adjList {
			if (cur.mask&(1<<v2)) > 0 || adjMatrix[cur.v][v2] < 1 || adjList[v2].rate < 1 {
				continue
			}

			nextState := cur
			nextState.t += adjMatrix[cur.v][v2] + 1
			if nextState.t >= limit {
				continue
			}
			nextState.mask |= 1 << v2
			nextState.v = v2
			nextState.flowed += nextState.rate * (adjMatrix[cur.v][v2] + 1)
			nextState.rate += adjList[v2].rate

			dpVal := nextState.flowed + (nextState.rate * (limit - nextState.t))

			if dp[nextState.mask] > dpVal {
				continue
			}
			dp[nextState.mask] = dpVal

			s = append(s, nextState)
		}

	}

	return fmt.Sprint(best)
}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)").Split(input, -1)

	rex := regexp.MustCompile(`Valve ([^\s]+) has flow rate=(\d+); tunnels? leads? to valves? ((?:[^,]+,?)+)`)

	adjList := []vertex{}
	vertexNameToId := map[string]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := rex.FindStringSubmatch(line)

		v0, rate, valves := matches[1], matches[2], strings.Split(matches[3], ", ")

		rateVal, err := strconv.Atoi(rate)
		if err != nil {
			log.Panic("Failed to read rate for line", line, err)
		}

		vertexNameToId[v0] = len(vertexNameToId)
		adjList = append(adjList, vertex{
			rateVal, valves,
		})
	}

	adjMatrix := make([][]int, len(adjList))

	for vStart := range len(adjList) {
		adjMatrix[vStart] = make([]int, len(adjList))

		//[mask, v]
		dist := 0
		q := [][2]int{{1 << vStart, vStart}}
		q2 := [][2]int{}

		for len(q) > 0 {
			cur := q[len(q)-1]
			q = q[:len(q)-1]
			mask, v := cur[0], cur[1]

			for _, v2s := range adjList[v].adj {
				v2 := vertexNameToId[v2s]

				if v2 == vStart || (adjMatrix[vStart][v2] > 0 && adjMatrix[vStart][v2] <= dist) {
					continue
				}

				adjMatrix[vStart][v2] = dist + 1
				mask2 := mask | (1 << v2)
				q2 = append(q2, [2]int{mask2, v2})
			}

			if len(q) < 1 {
				dist += 1
				q, q2 = q2, q
			}
		}
	}

	s := [][2]state{}

	{
		vStart := vertexNameToId["AA"]
		s = append(s, [2]state{{
			1 << vStart,
			vStart, //TODO this only kinda works because v0 has rate 0
			0,
			0,
			0,
		}, {
			1 << vStart,
			vStart, //TODO this only kinda works because v0 has rate 0
			0,
			0,
			0,
		}})
	}

	limit := 26
	best := 0

	adjToVisit := []int{}
	for i := range len(adjList) {
		if adjList[i].rate > 0 {
			adjToVisit = append(adjToVisit, i)
		}
	}

	dp := map[int64]int{}

	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]

		a, b := cur[0], cur[1]

		best = max(best,
			a.flowed+((limit-a.t)*a.rate)+
				b.flowed+((limit-b.t)*b.rate),
		)

		for _, v2 := range adjToVisit {
			if ((b.mask|a.mask)&(1<<v2)) > 0 || adjMatrix[a.v][v2] < 1 || adjList[v2].rate < 1 {
				continue
			}

			nextState := a
			nextState.t += adjMatrix[a.v][v2] + 1
			if nextState.t >= limit {
				continue
			}
			nextState.mask |= 1 << v2
			nextState.v = v2
			nextState.flowed += nextState.rate * (adjMatrix[a.v][v2] + 1)
			nextState.rate += adjList[v2].rate

			dpIdx := b.mask | nextState.mask
			dpVal := nextState.flowed + (nextState.rate * (limit - nextState.t)) +
				b.flowed + (b.rate * (limit - b.t))

			if dp[dpIdx] >= dpVal {
				continue
			}
			dp[nextState.mask] = dpVal

			s = append(s, [2]state{nextState, b})
		}

		for _, v2 := range adjToVisit {
			if ((b.mask|a.mask)&(1<<v2)) > 0 || adjMatrix[b.v][v2] < 1 || adjList[v2].rate < 1 {
				continue
			}

			nextState := b
			nextState.t += adjMatrix[b.v][v2] + 1
			if nextState.t >= limit {
				continue
			}
			nextState.mask |= 1 << v2
			nextState.v = v2
			nextState.flowed += nextState.rate * (adjMatrix[b.v][v2] + 1)
			nextState.rate += adjList[v2].rate

			dpIdx := a.mask | nextState.mask
			dpVal :=
				nextState.flowed + (nextState.rate * (limit - nextState.t)) +
					a.flowed + (a.rate * (limit - a.t))

			if dp[dpIdx] >= dpVal {
				continue
			}
			dp[dpIdx] = dpVal

			s = append(s, [2]state{a, nextState})
		}

	}

	return fmt.Sprint(best)
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "16")
	if err != nil {
		log.Fatal(err)
	}

	// input = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	// Valve BB has flow rate=13; tunnels lead to valves CC, AA
	// Valve CC has flow rate=2; tunnels lead to valves DD, BB
	// Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
	// Valve EE has flow rate=3; tunnels lead to valves FF, DD
	// Valve FF has flow rate=0; tunnels lead to valves EE, GG
	// Valve GG has flow rate=0; tunnels lead to valves FF, HH
	// Valve HH has flow rate=22; tunnel leads to valve GG
	// Valve II has flow rate=0; tunnels lead to valves AA, JJ
	// Valve JJ has flow rate=21; tunnel leads to valve II`

	fmt.Println(part1(input))

	fmt.Println(part2(input))

	log.Println((time.Since(start)))
}
