package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
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

	mappings := map[byte]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}
	mappingsRev := map[int]byte{}
	for k, v := range mappings {
		mappingsRev[v] = k
	}
	sum := 0

	for _, line := range lines {
		pow := 1
		num := 0
		for i := range len(line) {
			num += pow * mappings[line[(len(line)-1)-i]]
			pow *= 5
		}
		sum += num
	}

	convertToStr := func(n int) string {
		carry := 0
		val := []int{}
		for n > 0 || carry > 0 {
			rem := carry + (n % 5)
			carry = 0
			if rem == 0 {
				val = append(val, 0)
			} else if rem == 1 {
				val = append(val, 1)
			} else if rem == 2 {
				val = append(val, 2)
			} else if rem == 3 {
				carry = 1
				val = append(val, -2)
			} else if rem == 4 {
				carry = 1
				val = append(val, -1)
			} else {
				carry = 1
				val = append(val, 0)
			}
			n /= 5
		}

		strB := []byte{}
		for i := len(val) - 1; i >= 0; i-- {
			if len(strB) > 0 || val[i] != 0 {
				strB = append(strB, mappingsRev[val[i]])
			}
		}
		str := string(strB)

		return str
	}

	return convertToStr(sum)
}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "25")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))

	log.Println((time.Since(start)))
}
