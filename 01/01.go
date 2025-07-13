package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

func part1(input string) int {
	maxVal := 0
	total := 0

	for _, line := range regexp.MustCompile("\r*\n").Split(input, -1) {
		if strings.TrimSpace(line) == "" {
			total = 0

			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(fmt.Errorf("input contained non-numeric value %w", err))
		}

		total += num

		maxVal = max(total, maxVal)
	}

	return max(total, maxVal)
}

func part2(input string) int {
	total := 0

	h := pq.NewWith(func(a, b interface{}) int { return a.(int) - b.(int) })

	q := func(val int) {
		if h.Size() < 3 {
			h.Enqueue(val)
		} else {
			hVal, ok := h.Peek()
			if !ok {
				log.Fatal("impossible - empty heap?")
			}

			if hVal.(int) < val {
				h.Dequeue()
				h.Enqueue(val)
			}
		}
	}

	for _, line := range regexp.MustCompile("\r*\n").Split(input, -1) {
		if strings.TrimSpace(line) == "" {
			q(total)

			total = 0

			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(fmt.Errorf("input contained non-numeric value %w", err))
		}

		total += num

	}
	q(total)

	ret := 0
	for _, val := range h.Values() {
		ret += val.(int)
	}

	return ret
}

func main() {
	input, err := aoc.GetInput("2022", "1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
