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

type validRange struct {
	min       int
	max       int
	fieldName string
}

func part1(lines []string) (int, error) {
	endOfRules := 0
	rules := map[string][]validRange{}
	for i, line := range lines {
		if line == "" {
			endOfRules = i
			break
		}
		parts := strings.Split(line, ":")
		key := parts[0]
		ranges := strings.Trim(parts[1], " ")
		ruleParts := strings.Split(ranges, " or ")
		var parsedRanged []validRange
		for _, r := range ruleParts {
			rParts := strings.Split(r, "-")
			min, _ := strconv.Atoi(rParts[0])
			max, _ := strconv.Atoi(rParts[1])
			parsedRanged = append(parsedRanged, validRange{
				min: min,
				max: max,
			})
		}
		rules[key] = parsedRanged

	}

	var allRules []validRange
	for _, ranges := range rules {
		allRules = append(allRules, ranges...)
	}

	invalidSum := 0
	for _, line := range lines[endOfRules+5:] {
		parts := strings.Split(line, ",")
		nums, _ := utils.ToIntSlice(parts)
		for _, n := range nums {
			matchesRules := false
			for _, r := range allRules {
				matchesRules = matchesRules || (n >= r.min && n <= r.max)
			}
			if !matchesRules {
				invalidSum += n
			}
		}

	}

	return invalidSum, nil
}

func part2(lines []string) (int, error) {
	endOfRules := 0
	rules := map[string][]validRange{}
	for i, line := range lines {
		if line == "" {
			endOfRules = i
			break
		}
		parts := strings.Split(line, ":")
		key := parts[0]
		ranges := strings.Trim(parts[1], " ")
		ruleParts := strings.Split(ranges, " or ")
		var parsedRanged []validRange
		for _, r := range ruleParts {
			rParts := strings.Split(r, "-")
			min, _ := strconv.Atoi(rParts[0])
			max, _ := strconv.Atoi(rParts[1])
			parsedRanged = append(parsedRanged, validRange{
				min:       min,
				max:       max,
				fieldName: key,
			})
		}
		rules[key] = parsedRanged

	}

	var allRules []validRange
	for _, ranges := range rules {
		allRules = append(allRules, ranges...)
	}


	var possibleRules [][][]string

	for _, line := range lines[endOfRules+5:] {
		parts := strings.Split(line, ",")
		nums, _ := utils.ToIntSlice(parts)
		isInvalid := false
		var possibleRulesForTicket [][]string
		for _, n := range nums {
			var possibleRulesForField []string
			matchesRules := false
			for _, r := range allRules {
				matchesRules = matchesRules || (n >= r.min && n <= r.max)
				if n >= r.min && n <= r.max {
					possibleRulesForField = append(possibleRulesForField, r.fieldName)
				}
			}
			if !matchesRules {
				isInvalid = true
				break
			}
			possibleRulesForTicket = append(possibleRulesForTicket, possibleRulesForField)
		}
		if !isInvalid {
			possibleRules = append(possibleRules, possibleRulesForTicket)
		}

	}

	fieldNames := map[int][]string{}
	for fieldIndex := range possibleRules[0] {
		counts := map[string]int{}
		for _, line := range possibleRules {
			for _, fieldName := range line[fieldIndex] {
				counts[fieldName] ++
			}
		}
		for f, c := range counts {
			if c == len(possibleRules) {
				fieldNames[fieldIndex] = append(fieldNames[fieldIndex], f)
			}
		}
	}

	finalFieldNames := make([]string, len(fieldNames))

	for found := true; found; {
		found = false
		for i, names := range fieldNames {
			if len(names) == 1 {
				found = true
				finalFieldNames[i] = names[0]
				removeFromAll(fieldNames, names[0])
			}
		}
	}

	var startingWithDep []int
	for i, rule := range finalFieldNames {
		if strings.HasPrefix(rule, "departure") {
			startingWithDep = append(startingWithDep, i)
		}
	}

	println(len(startingWithDep))
	myTicket, _ := utils.ToIntSlice(strings.Split(lines[endOfRules+2], ","))
	result := 1
	for _, i := range startingWithDep {
		result *= myTicket[i]
	}

	return result, nil
}

func removeFromAll(fieldNames map[int][]string, s string) {
	for i, i2 := range fieldNames {
		for ruleIndex, rule := range i2 {
			if rule == s {
				fieldNames[i] = append(i2[:ruleIndex], i2[ruleIndex+1:]...)
				break
			}
		}

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
