package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"slices"
)

type Monke struct {
	worryLevels    []int
	operation      func(int) int
	condition      int
	conditionTrue  int
	conditionFalse int
}

func part1(input string) string {
	monkeyStrings := regexp.MustCompile("(\r*\n){2}").Split(input, -1)

	monkeys := []*Monke{}

	for _, monkeyString := range monkeyStrings {
		lines := regexp.MustCompile("(\r*\n)").Split(monkeyString, -1)

		//don't need monkey number
		lines = lines[1:]

		//starting items
		startingItemsStr := lines[0][len("  Starting items: "):]
		worryLevelsS := strings.Split(startingItemsStr, ", ")
		worryLevels := make([]int, len(worryLevelsS))
		for i, level := range worryLevelsS {
			lvl, err := strconv.Atoi(level)
			worryLevels[i] = lvl

			if err != nil {
				log.Fatal("bad worry level", worryLevels[i], lines[0], lines)
			}
		}
		lines = lines[1:]

		//operation
		operation := lines[0][len("  Operation: new = "):]
		lines = lines[1:]

		//test
		test, err := strconv.Atoi(lines[0][len("  Test: divisible by "):])
		if err != nil {
			log.Fatal("bad test:", lines[0], err)
		}
		lines = lines[1:]
		testTrue, err := strconv.Atoi(lines[0][len("    If true: throw to monkey "):])
		if err != nil {
			log.Fatal("bad test true branch:", lines[0], err)
		}
		lines = lines[1:]
		testFalse, err := strconv.Atoi(lines[0][len("    If false: throw to monkey "):])
		if err != nil {
			log.Fatal("bad test false branch:", lines[0], err)
		}
		lines = lines[1:]

		monkeys = append(monkeys, &Monke{
			worryLevels,
			buildOperationFn(operation),
			test,
			testTrue,
			testFalse,
		})
	}

	inspectionCounts := make([]int, len(monkeys))

	for range 20 {
		for mIdx, monke := range monkeys {
			for _, worryLevel := range monke.worryLevels {
				worryLevel = monke.operation(worryLevel)
				worryLevel /= 3

				if (worryLevel % monke.condition) == 0 {
					//true
					monkeys[monke.conditionTrue].worryLevels = append(monkeys[monke.conditionTrue].worryLevels, worryLevel)
				} else {
					// false
					monkeys[monke.conditionFalse].worryLevels = append(monkeys[monke.conditionFalse].worryLevels, worryLevel)
				}
			}
			inspectionCounts[mIdx] += len(monke.worryLevels)
			monke.worryLevels = []int{}
		}
	}

	slices.Sort(inspectionCounts)

	return fmt.Sprint(inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2])

}

func buildOperationFn(opStr string) func(int) int {
	s := strings.Split(opStr, " ")

	operand1, op, operand2 := s[0], s[1], s[2]

	old1, old2 := operand1 == "old", operand2 == "old"
	val1, val2 := 0, 0

	if !old1 {
		v, err := strconv.Atoi(operand1)
		val1 = v
		if err != nil {
			log.Fatal("failed to parse expression - bad op1 ", operand1, opStr)
		}
	}
	if !old2 {
		v, err := strconv.Atoi(operand2)
		val2 = v
		if err != nil {
			log.Fatal("failed to parse expression - bad op2 ", operand2, opStr)
		}
	}

	return func(in int) int {
		left, right := val1, val2
		if old1 {
			left = in
		}
		if old2 {
			right = in
		}

		if op == "*" {
			return left * right
		}
		if op == "+" {
			return left + right
		}

		log.Fatal("Unknown operation ", op, opStr)
		panic("")
	}
}

func part2(input string) string {
	monkeyStrings := regexp.MustCompile("(\r*\n){2}").Split(input, -1)

	monkeys := []*Monke{}

	for _, monkeyString := range monkeyStrings {
		lines := regexp.MustCompile("(\r*\n)").Split(monkeyString, -1)

		//don't need monkey number
		lines = lines[1:]

		//starting items
		startingItemsStr := lines[0][len("  Starting items: "):]
		worryLevelsS := strings.Split(startingItemsStr, ", ")
		worryLevels := make([]int, len(worryLevelsS))
		for i, level := range worryLevelsS {
			lvl, err := strconv.Atoi(level)
			worryLevels[i] = lvl

			if err != nil {
				log.Fatal("bad worry level", worryLevels[i], lines[0], lines)
			}
		}
		lines = lines[1:]

		//operation
		operation := lines[0][len("  Operation: new = "):]
		lines = lines[1:]

		//test
		test, err := strconv.Atoi(lines[0][len("  Test: divisible by "):])
		if err != nil {
			log.Fatal("bad test:", lines[0], err)
		}
		lines = lines[1:]
		testTrue, err := strconv.Atoi(lines[0][len("    If true: throw to monkey "):])
		if err != nil {
			log.Fatal("bad test true branch:", lines[0], err)
		}
		lines = lines[1:]
		testFalse, err := strconv.Atoi(lines[0][len("    If false: throw to monkey "):])
		if err != nil {
			log.Fatal("bad test false branch:", lines[0], err)
		}
		lines = lines[1:]

		monkeys = append(monkeys, &Monke{
			worryLevels,
			buildOperationFn(operation),
			test,
			testTrue,
			testFalse,
		})
	}

	modVal := 1
	for _, monke := range monkeys {
		modVal *= monke.condition
	}

	inspectionCounts := make([]int, len(monkeys))

	for range 10000 {
		for mIdx, monke := range monkeys {
			for _, worryLevel := range monke.worryLevels {
				worryLevel = monke.operation(worryLevel) % modVal

				if (worryLevel % monke.condition) == 0 {
					//true
					monkeys[monke.conditionTrue].worryLevels = append(monkeys[monke.conditionTrue].worryLevels, worryLevel)
				} else {
					// false
					monkeys[monke.conditionFalse].worryLevels = append(monkeys[monke.conditionFalse].worryLevels, worryLevel)
				}
			}
			inspectionCounts[mIdx] += len(monke.worryLevels)
			monke.worryLevels = []int{}
		}
	}

	slices.Sort(inspectionCounts)

	return fmt.Sprint(inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2])
}

func main() {
	input, err := aoc.GetInput("2022", "11")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(part1(input))
	log.Println(part2(input))
}
