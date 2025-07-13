package main

import (
	"aoc2022/aoc"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type folder struct {
	id            int
	name          string
	depth         int
	parent        *folder
	children      map[string]*folder
	size          int
	childrenAdded int
}

func part1(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	folderId := 1

	curFolder := &folder{children: map[string]*folder{}}
	leaves := map[int]*folder{0: curFolder}
	root := curFolder

	for i := 0; i < len(lines); {
		if lines[i] == "" {
			i++
			continue
		}

		if lines[i][0] != '$' {
			log.Panic("Expected command")
		}

		line := lines[i]

		cmd := line[2:]

		if cmd[:2] == "ls" {
			i++
			for i < len(lines) {
				if lines[i] == "" || lines[i][0] == '$' {
					break
				}

				if len(lines[i]) >= 5 && lines[i][0:3] == "dir" {
					//disregard
					i++
					continue
				}

				split := strings.Split(lines[i], " ")

				size, err := strconv.Atoi(split[0])
				if err != nil {
					log.Panic("bad file size", lines[i])
				}

				curFolder.size += size

				i++
			}
			i--
		} else if cmd[:2] == "cd" {
			target := cmd[3:]

			if target == ".." {
				curFolder = curFolder.parent
			} else if target == "/" {
				curFolder = root
			} else {
				if _, ok := curFolder.children[target]; !ok {
					curFolder.children[target] = &folder{
						id:       folderId,
						name:     target,
						depth:    curFolder.depth + 1,
						parent:   curFolder,
						children: map[string]*folder{},
					}
					leaves[folderId] = curFolder.children[target]
					folderId += 1
				}

				delete(leaves, curFolder.id)

				curFolder = curFolder.children[target]
			}
		} else {
			log.Panic("Unexpected command", cmd)
		}

		i++
	}

	result := 0

	idsSeen := map[int]bool{}

	q := leaves
	q2 := map[int]*folder{}
	for len(q) > 0 {
		for _, curFolder := range q {
			if idsSeen[curFolder.id] {
				log.Panic("broken")
			}
			idsSeen[curFolder.id] = true

			if curFolder.size <= 100000 {
				result += curFolder.size
			}

			if curFolder.parent == nil {
				continue
			}

			curFolder.parent.size += curFolder.size
			curFolder.parent.childrenAdded += 1
			if curFolder.parent.childrenAdded == len(curFolder.parent.children) {
				q2[curFolder.parent.id] = curFolder.parent
			}
		}

		q = q2
		q2 = map[int]*folder{}
	}

	return fmt.Sprintln(result)
}

func part2(input string) string {
	lines := regexp.MustCompile("(\r*\n)+").Split(input, -1)

	folderId := 1

	curFolder := &folder{children: map[string]*folder{}}
	leaves := map[int]*folder{0: curFolder}
	root := curFolder

	for i := 0; i < len(lines); {
		if lines[i] == "" {
			i++
			continue
		}

		if lines[i][0] != '$' {
			log.Panic("Expected command")
		}

		line := lines[i]

		cmd := line[2:]

		if cmd[:2] == "ls" {
			i++
			for i < len(lines) {
				if lines[i] == "" || lines[i][0] == '$' {
					break
				}

				if len(lines[i]) >= 5 && lines[i][0:3] == "dir" {
					//disregard
					i++
					continue
				}

				split := strings.Split(lines[i], " ")

				size, err := strconv.Atoi(split[0])
				if err != nil {
					log.Panic("bad file size", lines[i])
				}

				curFolder.size += size

				i++
			}
			i--
		} else if cmd[:2] == "cd" {
			target := cmd[3:]

			if target == ".." {
				curFolder = curFolder.parent
			} else if target == "/" {
				curFolder = root
			} else {
				if _, ok := curFolder.children[target]; !ok {
					curFolder.children[target] = &folder{
						id:       folderId,
						name:     target,
						depth:    curFolder.depth + 1,
						parent:   curFolder,
						children: map[string]*folder{},
					}
					leaves[folderId] = curFolder.children[target]
					folderId += 1
				}

				delete(leaves, curFolder.id)

				curFolder = curFolder.children[target]
			}
		} else {
			log.Panic("Unexpected command", cmd)
		}

		i++
	}

	idsSeen := map[int]bool{}

	q := leaves
	q2 := map[int]*folder{}
	for len(q) > 0 {
		for _, curFolder := range q {
			if idsSeen[curFolder.id] {
				log.Panic("broken")
			}
			idsSeen[curFolder.id] = true

			if curFolder.parent == nil {
				continue
			}

			curFolder.parent.size += curFolder.size
			curFolder.parent.childrenAdded += 1
			if curFolder.parent.childrenAdded == len(curFolder.parent.children) {
				q2[curFolder.parent.id] = curFolder.parent
			}
		}

		q = q2
		q2 = map[int]*folder{}
	}

	total := 70000000
	needed := 30000000
	used := root.size
	toFree := needed - (total - used)
	bestFit := root.size
	{
		s := []*folder{root}

		for len(s) > 0 {
			curFolder := s[len(s)-1]
			s = s[:len(s)-1]

			if curFolder.size < bestFit && curFolder.size >= toFree {
				bestFit = curFolder.size
			}

			for _, child := range curFolder.children {
				s = append(s, child)
			}
		}
	}

	return fmt.Sprintln(bestFit)
}

func main() {
	input, err := aoc.GetInput("2022", "7")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
