package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type NumberStats struct {
	lastTimeSpoken int
	count          int
	nextNumber     int
}

func part1(lines []string) (int, error) {

	return calcNthNumber(lines, 2020), nil
}

func calcNthNumber(lines []string, n int) int {
	mem := map[int]NumberStats{}

	numbers, _ := utils.ToIntSlice(lines)
	for i, e := range numbers {
		mem[e] = NumberStats{
			lastTimeSpoken: i + 1,
			count:          0,
			nextNumber:     0,
		}
	}
	spoken := numbers[len(numbers)-1]

	for i := len(numbers) + 1; i <= n; i++ {
		stats, _ := mem[spoken]

		spoken = stats.nextNumber
		stats, exists := mem[spoken]
		mem[spoken] = updateStats(stats, exists, i)
	}

	return spoken
}

func updateStats(stats NumberStats, isKnown bool, lastTime int) NumberStats {
	if isKnown {
		return NumberStats{
			lastTimeSpoken: lastTime,
			count:          stats.count + 1,
			nextNumber:     lastTime - stats.lastTimeSpoken,
		}
	} else {
		return NumberStats{
			lastTimeSpoken: lastTime,
			count:          0,
			nextNumber:     0,
		}
	}

}

func part2(lines []string) (int, error) {
	return calcNthNumber(lines, 30000000), nil
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
