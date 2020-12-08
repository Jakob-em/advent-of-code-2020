package main

import (
	"errors"
	"fmt"
	"github.com/Jakob-em/advent-of-code-2020/simulator"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type Instruction struct {
	opcode   string
	operator int
}

func part1(lines []string) (int, error) {
	sim, err := simulator.New(lines, nil)

	if err != nil {
		return 0, err
	}
	sim.Run()

	return sim.Accumulator, nil

}

func part2(lines []string) (int, error) {

	for _, lineNumToChange := range findPossibleChanges(lines) {

		sim, err := simulator.New(lines, nil)

		if err != nil {
			return 0, err
		}

		switchNopWithJmp(&sim.Code[lineNumToChange])

		sim.Run()
		if sim.IsAtEnd() {
			return sim.Accumulator, nil
		}
	}

	return 0, errors.New("no line change found that has no infinite loop")
}

func switchNopWithJmp(instruction *simulator.Instruction) {
	if instruction.Operation == "jmp" {
		instruction.Operation = "nop"
	} else if instruction.Operation == "nop" {
		instruction.Operation = "jmp"

	}
}

func findPossibleChanges(lines []string) []int {
	var nums []int
	for i, line := range lines {
		if strings.Contains(line, "jmp") || strings.Contains(line, "nop") {
			nums = append(nums, i)
		}
	}
	return nums
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
