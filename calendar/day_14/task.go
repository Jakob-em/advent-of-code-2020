package main

import (
	"fmt"
	"github.com/Jakob-em/advent-of-code-2020/utils"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func part1(lines []string) (int64, error) {

	var mask string
	memory := map[int64]int64{}
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[7:]
		} else {
			address, value := getValueAndAddress(line)
			masked := applyMask(mask, value)
			memory[int64(address)] = masked
		}
	}

	return sumValuesInMemory(memory), nil
}

func getValueAndAddress(line string) (address int, value int) {
	groups := utils.ExtractGroups(`\[(?P<address>\d+)\] = (?P<value>\d+)`, line)
	address, _ = strconv.Atoi(groups["address"])
	value, _ = strconv.Atoi(groups["value"])
	return address, value
}

func applyMask(mask string, val int) int64 {
	binary := fmt.Sprintf("%036b", val)
	withMask := ""
	for i, m := range mask {
		if m == 'X' {
			withMask += string(binary[i])
		} else {
			withMask += string(m)
		}
	}
	maskedValue, _ := strconv.ParseInt(withMask, 2, 64)
	return maskedValue
}

func part2(lines []string) (int64, error) {
	var mask string
	memory := map[int64]int64{}
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[7:]
		} else {
			address, value := getValueAndAddress(line)
			for _, a := range getPossibleFloatingValues(mask, address) {
				memory[a] = int64(value)
			}
		}
	}
	sum := sumValuesInMemory(memory)
	return sum, nil
}

func sumValuesInMemory(memory map[int64]int64) int64 {
	sum := int64(0)
	for _, m := range memory {
		sum += m
	}
	return sum
}

func getPossibleFloatingValues(mask string, address int) []int64 {
	binary := fmt.Sprintf("%036b", address)
	withMask := ""
	for i, m := range mask {
		if m == '0' {
			withMask += string(binary[i])
		} else {
			withMask += string(m)
		}
	}

	return getFloatingValuesRec(withMask)
}

func getFloatingValuesRec(floating string) []int64 {
	if strings.Contains(floating, "X") {
		withOne := getFloatingValuesRec(strings.Replace(floating, "X", "1", 1))
		withZero := getFloatingValuesRec(strings.Replace(floating, "X", "0", 1))
		return append(withOne, withZero...)
	} else {
		val, _ := strconv.ParseInt(floating, 2, 64)
		return []int64{val}
	}
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
