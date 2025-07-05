package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
	"time"
)

func part1(input string) string {
	pairs := regexp.MustCompile("(\r*\n){2}").Split(input, -1)

	ret := 0

	for idx, pairS := range pairs {
		pairSplit := strings.Split(pairS, "\n")

		a, b := pairSplit[0], pairSplit[1]

		c := cmp(a, b)
		if c < 0 {
			ret += idx + 1
		}

	}

	return fmt.Sprint(ret)
}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sep1 := "[[2]]"
	sep2 := "[[6]]"
	lines = append(lines, sep1, sep2)

	slices.SortFunc(lines, cmp)

	ret := 1
	for idx, line := range lines {
		if line == sep1 {
			ret *= idx + 1
		}
		if line == sep2 {
			ret *= idx + 1
		}
	}

	return fmt.Sprint(ret)
}

func cmp(a, b string) int {

	aIdx, bIdx := 0, 0

	readInt := func(s string, idx int) (int, int) {
		val := 0

		for idx < len(s) {
			if s[idx] <= '9' && s[idx] >= '0' {
				val *= 10
				val += int(s[idx] - '0')
				idx++
			} else {
				break
			}
		}

		return idx, val
	}

	var cmp func() int
	cmp = func() int {
		aInt := a[aIdx] >= '0' && a[aIdx] <= '9'
		bInt := b[bIdx] >= '0' && b[bIdx] <= '9'
		if aInt && bInt {
			aVal, bVal := 0, 0
			aIdx, aVal = readInt(a, aIdx)
			bIdx, bVal = readInt(b, bIdx)

			return aVal - bVal
		}

		if a[aIdx] == '[' && b[bIdx] == '[' {
			aIdx += 1 //[
			bIdx += 1

			for {
				c := cmp()
				if c > 0 {
					return 1
				}
				if c < 0 {
					return -1
				}

				if a[aIdx] == ',' && b[bIdx] == ',' {
					aIdx++
					bIdx++
				} else if a[aIdx] == ']' && b[bIdx] == ']' {
					break
				} else {
					if a[aIdx] == ']' {
						//b too long
						return -1
					} else if b[bIdx] == ']' {
						//a too long
						return 1
					}
				}
			}

			aIdx += 1 //]
			bIdx += 1

			return 0
		}

		if a[aIdx] == ']' && b[bIdx] == ']' {
			return 0
		}
		if a[aIdx] == ']' {
			return -1
		} else if b[bIdx] == ']' {
			return 1
		}

		if a[aIdx] == '[' {
			aIdx++

			ret := cmp()

			if ret == 0 && a[aIdx] != ']' {
				return 1
			} else {
				aIdx++
			}

			return ret
		} else if b[bIdx] == '[' {
			//treat a as [a]
			bIdx++
			ret := cmp()

			if ret == 0 && b[bIdx] != ']' {
				return -1
			} else {
				bIdx++
			}

			return ret
		}

		panic("wrong")

	}

	return cmp()

}

func main() {
	start := time.Now()

	input, err := aoc.GetInput("2022", "13")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(part1(input))
	log.Println(part2(input))

	log.Println((time.Since(start)))
}
