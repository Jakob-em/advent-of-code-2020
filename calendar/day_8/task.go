package main

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type Instruction struct {
	opcode   string
	operator int
}

func part1(linesOrg []string) (int, error) {
	accumulator := 0
	lines := copyLines(linesOrg)

	for i := 0; ; {
		if lines[i] == "-" {
			break
		}
		instruction, err := parseInstruction(lines[i])
		lines[i] = "-"

		if err != nil {
			return 0, err
		}
		accumulator, i = executeInstruction(instruction, accumulator, i)

	}

	return accumulator, nil
}

func executeInstruction(instruction Instruction, accumulator int, i int) (newAccumulator int, newIndex int) {
	if instruction.opcode == "acc" {
		return accumulator + instruction.operator, i + 1
	}

	if instruction.opcode == "jmp" {
		return accumulator, i + instruction.operator

	}

	return accumulator, i + 1
}

func copyLines(linesOrg []string) []string {
	lines := make([]string, len(linesOrg))
	copy(lines, linesOrg)
	return lines
}

func part2(linesOrg []string) (int, error) {

	for _, lineNumToChange := range findPossibleChanges(linesOrg) {
		i := 0
		accumulator := 0
		lines := copyLines(linesOrg)

		for i < len(lines) {
			if lines[i] == "-" {
				break
			}

			if i == lineNumToChange {
				lines[i] = replaceLine(lines[i])
			}

			instruction, err := parseInstruction(lines[i])
			lines[i] = "-"

			if err != nil {
				return 0, err
			}
			accumulator, i = executeInstruction(instruction, accumulator, i)

		}
		if i == len(lines) {
			return accumulator, nil
		}
	}

	return 0, errors.New("no line change found that has no infinite loop")
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

func replaceLine(line string) string {
	if strings.Contains(line, "jmp") {
		return strings.Replace(line, "jmp", "nop", -1)
	}
	return strings.Replace(line, "nop", "jmp", -1)
}

func parseInstruction(line string) (Instruction, error) {
	cmd := strings.Split(line, " ")
	arg, err := strconv.Atoi(cmd[1])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{
		opcode:   cmd[0],
		operator: arg,
	}, nil
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
