package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type op struct {
	val        int
	v1, v2, op string
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nodes := map[string]op{}

	for _, line := range lines {
		split := strings.Split(line, ": ")
		v0 := split[0]

		splitOp := strings.Split(split[1], " ")
		if len(splitOp) == 1 {
			num, err := strconv.Atoi(splitOp[0])
			if err != nil {
				log.Panic("failed to parse line", line, err)
			}
			nodes[v0] = op{val: num}
		} else {
			v1 := splitOp[0]
			v2 := splitOp[2]
			operation := splitOp[1]
			nodes[v0] = op{0, v1, v2, operation}
		}
	}

	var dfs func(node string) int
	dfs = func(node string) int {
		v0 := nodes[node]
		if v0.v1 == "" {
			return v0.val
		}

		switch v0.op {
		case "+":
			return dfs(v0.v1) + dfs(v0.v2)
		case "-":
			return dfs(v0.v1) - dfs(v0.v2)
		case "*":
			return dfs(v0.v1) * dfs(v0.v2)
		case "/":
			return dfs(v0.v1) / dfs(v0.v2)
		default:
			log.Panic("Unknown op")
		}

		return -1
	}

	return fmt.Sprint(dfs("root"))
}

func part2(input string) string {
	lines := strings.Split(input, "\n")

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nodes := map[string]*op{}
	nodeParents := map[string]string{}

	for _, line := range lines {
		split := strings.Split(line, ": ")
		v0 := split[0]

		splitOp := strings.Split(split[1], " ")
		if len(splitOp) == 1 {
			num, err := strconv.Atoi(splitOp[0])
			if err != nil {
				log.Panic("failed to parse line", line, err)
			}
			nodes[v0] = &op{val: num}
		} else {
			v1 := splitOp[0]
			v2 := splitOp[2]
			operation := splitOp[1]
			nodes[v0] = &op{0, v1, v2, operation}

			nodeParents[v1] = v0
			nodeParents[v2] = v0
		}
	}

	var dfs func(node string) int
	dfs = func(node string) int {
		v0 := nodes[node]
		if v0.v1 == "" {
			return v0.val
		}

		switch v0.op {
		case "+":
			return dfs(v0.v1) + dfs(v0.v2)
		case "-":
			return dfs(v0.v1) - dfs(v0.v2)
		case "*":
			return dfs(v0.v1) * dfs(v0.v2)
		case "/":
			return dfs(v0.v1) / dfs(v0.v2)
		default:
			log.Panic("Unknown op")
		}

		return -1
	}

	var dfsUp func(childNode string) int
	dfsUp = func(childNode string) int {
		parent := nodeParents[childNode]
		parentNode := nodes[parent]

		switch parentNode.op {
		case "+":

			if parentNode.v1 == childNode {
				return dfsUp(parent) - dfs(parentNode.v2)
			} else {
				return dfsUp(parent) - dfs(parentNode.v1)
			}

		case "-":

			if parentNode.v1 == childNode {
				return dfsUp(parent) + dfs(parentNode.v2)
			} else {
				return dfs(parentNode.v1) - dfsUp(parent)
			}

		case "*":

			if parentNode.v1 == childNode {
				return dfsUp(parent) / dfs(parentNode.v2)
			} else {
				return dfsUp(parent) / dfs(parentNode.v1)
			}

		case "/":

			if parentNode.v1 == childNode {
				return dfsUp(parent) * dfs(parentNode.v2)
			} else {
				return dfs(parentNode.v1) / dfsUp(parent)
			}

		case "=":

			if parentNode.v1 == childNode {
				return dfs(parentNode.v2)
			} else {
				return dfs(parentNode.v1)
			}

		default:
			log.Panic("Unknown op")
		}
		panic("No")

	}

	nodes["root"].op = "="

	return fmt.Sprint(dfsUp("humn"))
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "21")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))

	fmt.Println((time.Since(start)))
}
