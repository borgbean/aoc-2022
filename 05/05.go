package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func part1(input string) string {
	parts := regexp.MustCompile("(\r*\n){2}").Split(input, 2)
	crateLines := regexp.MustCompile("(\r*\n)+").Split(parts[0], -1)

	columnsCount := (len(crateLines[0]) + 1) / 4
	columns := make([][]byte, columnsCount)

	for i := len(crateLines) - 2; i >= 0; i-- {
		for col := 0; col < columnsCount; col++ {
			chr := crateLines[i][1+(4*col)]

			if chr != ' ' {
				columns[col] = append(columns[col], chr)
			}
		}
	}

	movementLines := regexp.MustCompile("(\r*\n)+").Split(parts[1], -1)

	rex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, line := range movementLines {
		if line == "" {
			continue
		}
		split := rex.FindStringSubmatch(line)

		count, err1 := strconv.Atoi(split[1])
		from, err2 := strconv.Atoi(split[2])
		to, err3 := strconv.Atoi(split[3])

		if err1 != nil || err2 != nil || err3 != nil {
			log.Panic(err1, err2, err3)
		}

		from -= 1
		to -= 1

		for i := range count {
			columns[to] = append(columns[to], columns[from][(len(columns[from])-1)-i])
		}
		columns[from] = columns[from][:len(columns[from])-count]
	}

	ret := []byte{}
	for col := range columns {
		if len(columns[col]) < 1 {
			continue
		}

		ret = append(ret, columns[col][len(columns[col])-1])
	}

	return string(ret)
}

func part2(input string) string {

	parts := regexp.MustCompile("(\r*\n){2}").Split(input, 2)
	crateLines := regexp.MustCompile("(\r*\n)+").Split(parts[0], -1)

	columnsCount := (len(crateLines[0]) + 1) / 4
	columns := make([][]byte, columnsCount)

	for i := len(crateLines) - 2; i >= 0; i-- {
		for col := 0; col < columnsCount; col++ {
			chr := crateLines[i][1+(4*col)]

			if chr != ' ' {
				columns[col] = append(columns[col], chr)
			}
		}
	}

	movementLines := regexp.MustCompile("(\r*\n)+").Split(parts[1], -1)

	rex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, line := range movementLines {
		if line == "" {
			continue
		}
		split := rex.FindStringSubmatch(line)

		count, err1 := strconv.Atoi(split[1])
		from, err2 := strconv.Atoi(split[2])
		to, err3 := strconv.Atoi(split[3])

		if err1 != nil || err2 != nil || err3 != nil {
			log.Panic(err1, err2, err3)
		}

		from -= 1
		to -= 1

		columns[to] = append(columns[to], columns[from][len(columns[from])-count:]...)
		columns[from] = columns[from][:len(columns[from])-count]
	}

	ret := []byte{}
	for col := range columns {
		if len(columns[col]) < 1 {
			continue
		}

		ret = append(ret, columns[col][len(columns[col])-1])
	}

	return string(ret)
}

func main() {
	input, err := aoc.GetInput("2022", "5")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
