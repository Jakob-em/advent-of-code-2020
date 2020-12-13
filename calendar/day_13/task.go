package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type busInfo struct {
	id            int
	minTimeToWait int
}

func part1(lines []string) (int, error) {
	earliestDeparture, _ := strconv.Atoi(lines[0])
	var min busInfo
	for _, s := range strings.Split(lines[1], ",") {
		if s != "x" {
			id, _ := strconv.Atoi(s)
			waitTime := id - (earliestDeparture % id)
			if min.minTimeToWait == 0 || waitTime < min.minTimeToWait {
				min = busInfo{
					id:            id,
					minTimeToWait: waitTime,
				}
			}
		}
	}

	return min.id * min.minTimeToWait, nil
}

type congruence struct {
	mod       int
	remainder int
}

func solveSystem(congruences []congruence) int {
	sum := 0
	m := 1
	ch := make(chan int, len(congruences))
	for _, c := range congruences {
		m *= c.mod
	}
	for _, c := range congruences {
		go func(c congruence) {
			M := m / c.mod
			N := multiplicativeInverse(M, c.mod)
			ch <- M * N * c.remainder
		}(c)
	}

	for range congruences {
		sum += <-ch
	}

	return sum % m
}

func multiplicativeInverse(a int, m int) int {
	b := m
	x0, x1 := 1, 0
	for b != 0 {
		newB := a % b
		newX := x0 - x1*(a/b)
		x0 = x1
		x1 = newX
		a = b
		b = newB
	}

	return (x0 + m) % m
}

func part2(lines []string) (int, error) {
	var congruences []congruence

	for i, s := range strings.Split(lines[1], ",") {
		if s != "x" {
			id, _ := strconv.Atoi(s)
			congruences = append(congruences, congruence{
				mod:       id,
				remainder: ((id * 2) - i) % id,
			})
		}
	}

	return solveSystem(congruences), nil
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
