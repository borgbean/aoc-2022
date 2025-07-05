package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type node struct {
	prev, next *node
	val        int64
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nodes := make([]*node, 0, len(lines))

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Panic("failed to parse line", line, err)
		}

		nodes = append(nodes, &node{
			nil, nil, int64(num),
		})

		if len(nodes) > 1 {
			nodes[len(nodes)-1].prev = nodes[len(nodes)-2]
			nodes[len(nodes)-2].next = nodes[len(nodes)-1]
		}
	}
	nodes[0].prev = nodes[len(nodes)-1]
	nodes[len(nodes)-1].next = nodes[0]

	for _, node := range nodes {
		cur := node
		if cur.val == 0 {
			continue
		}
		cur.prev.next = cur.next
		cur.next.prev = cur.prev
		for range max(node.val, -node.val) {
			if node.val > 0 {
				//forward
				cur = cur.next
			} else {
				cur = cur.prev
			}
		}
		if node.val > 0 {
			node.next = cur.next
			node.prev = cur
			cur.next = node
			node.next.prev = node
		} else {
			node.next = cur
			node.prev = cur.prev
			cur.prev = node
			node.prev.next = node
		}

	}

	sum := int64(0)
	{
		cur := nodes[0]

		for cur.val != 0 {
			cur = cur.next
		}

		for i := 1; i < 3001; i++ {
			cur = cur.next
			if (i % 1000) == 0 {
				sum += cur.val
			}
		}
	}

	return fmt.Sprint(sum)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nodes := make([]*node, 0, len(lines))

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Panic("failed to parse line", line, err)
		}

		nodes = append(nodes, &node{
			nil, nil, int64(num) * 811589153,
		})

		if len(nodes) > 1 {
			nodes[len(nodes)-1].prev = nodes[len(nodes)-2]
			nodes[len(nodes)-2].next = nodes[len(nodes)-1]
		}
	}
	nodes[0].prev = nodes[len(nodes)-1]
	nodes[len(nodes)-1].next = nodes[0]

	for range 10 {
		for _, node := range nodes {
			cur := node

			val := int(max(node.val, -node.val) % int64(len(nodes)-1))

			if val == 0 {
				continue
			}

			positive := node.val > 0
			if val > (len(nodes) / 2) {
				val = (len(nodes) - 1) - val
				positive = !positive
			}

			cur.prev.next = cur.next
			cur.next.prev = cur.prev

			for range val {
				if positive {
					//forward
					cur = cur.next
				} else {
					cur = cur.prev
				}
			}
			if positive {
				node.next = cur.next
				node.prev = cur
				cur.next = node
				node.next.prev = node
			} else {
				node.next = cur
				node.prev = cur.prev
				cur.prev = node
				node.prev.next = node
			}

		}
	}

	sum := int64(0)
	{
		cur := nodes[0]

		for cur.val != 0 {
			cur = cur.next
		}

		for i := 1; i < 3001; i++ {
			cur = cur.next
			if (i % 1000) == 0 {
				sum += cur.val
			}
		}
	}

	return fmt.Sprint(sum)
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "20")
	if err != nil {
		log.Fatal(err)
	}

	// 	input = `1
	// 2
	// -3
	// 3
	// -2
	// 0
	// 4
	// `

	fmt.Println(part1(input))

	fmt.Println(part2(input))

	fmt.Println((time.Since(start)))
}
