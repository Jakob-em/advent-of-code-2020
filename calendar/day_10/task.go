package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

func part1(lines []string) (int, error) {

	withDiffOne := 0
	withDiffThree := 0

	lastMin := 0
	remaining, _ := utils.ToIntSlice(lines)
	remaining = append(remaining, utils.Max(remaining)+3)
	for ; ; {
		if len(remaining) == 0 {
			break
		}
		min, index := utils.MinWithIndex(remaining)
		if lastMin == min-1 {
			withDiffOne += 1
		} else if lastMin == min-3 {
			withDiffThree += 1
		}
		remaining = append(remaining[:index], remaining[index+1:]...)
		lastMin = min
	}

	return withDiffOne * withDiffThree, nil
}

func part2(lines []string) (int, error) {
	ints, _ := utils.ToIntSlice(lines)

	rated := utils.Max(ints)
	rated += 3

	connections := map[int][]int{}
	connections[0] = []int{1, 2, 3}
	for _, i := range ints {
		connections[i] = []int{i + 1, i + 2, i + 3}
	}

	return countConnectionsToDestination(0, connections, rated, map[int]int{}), nil
}

func countConnectionsToDestination(node int, connections map[int][]int, dest int, cache map[int]int) int {
	if connections[node] == nil {
		return 0
	}

	c, cached := cache[node]
	if cached {
		return c
	}
	count := 0
	for _, c := range connections[node] {
		if c == dest {
			count += 1
		} else {
			count += countConnectionsToDestination(c, connections, dest, cache)
		}
	}
	cache[node] = count
	return count
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	inputFile := filepath.Join(path.Dir(filename), "input.txt")
	lines := utils.ReadLinesFromFile(inputFile, "\n")

	result, err := part1(lines)
	if err != nil {
		log.Fatalf("Part 1: %s", err)
	}
	fmt.Printf("Result from Part 1: %d\n", result)

	result, err = part2(lines)
	if err != nil {
		log.Fatalf("Part 2: %s", err)
	}
	fmt.Printf("Result from Part 2: %d\n", result)
}
