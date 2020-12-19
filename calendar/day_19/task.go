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

//var depth = 0

//func (c *ComposedRule) buildRegex(rules map[int]Rule) string {
//	depth++
//	println(depth)
//	if depth > 5000 {
//		return "(xxx)"
//	}
//	if c.cache != "" {
//		return c.cache
//	}
//	result := ""
//	for i, chainedRules := range c.rules {
//		var partRule = "("
//		for _, rule := range chainedRules {
//			partRule += fmt.Sprintf("(%s)", rules[rule].buildRegex(rules))
//		}
//		result += partRule + ")"
//		if i != len(c.rules)-1 {
//			result += "|"
//		}
//	}
//	c.cache = result
//	depth--
//	return result
//}

//
//func (v *ValueRule) buildRegex(rules map[int]Rule) string {
//	return v.value
//}

type Rule interface {
	match(value string, rules RuleLookup) []result
}

type ComposedRule struct {
	rules [][]int
}

type RuleLookup = *map[int]Rule

type result struct {
	remaining string
	matched   bool
}

var depth = 0

func (c *ComposedRule) match(value string, rules RuleLookup) []result {
	depth++
	if depth > 100 {
		return []result{}
	}
	var allResults []result
	for _, chainedRule := range c.rules {
		matchedAll := true
		toCheck := []result{{
			remaining: value,
			matched:   true,
		}}
		for _, ruleId := range chainedRule {
			var found []result
			for _, r := range toCheck {
				matches := (*rules)[ruleId].match(r.remaining, rules)
				found = append(found, matches...)
			}

			if len(found) == 0 {
				matchedAll = false
				break
			}
			toCheck = found
		}
		if matchedAll {
			allResults = append(allResults, toCheck...)
		}
	}
	return allResults
}

type ValueRule struct {
	value string
}

func (v *ValueRule) match(value string, rules RuleLookup) []result {
	return []result{
		{
			remaining: strings.Replace(value, v.value, "", 1),
			matched:   strings.HasPrefix(value, v.value),
		},
	}
}

func part1(lines []string) (int, error) {
	rules := map[int]Rule{}
	lastRule := 0
	for i, line := range lines {
		if line == "" {
			break
		}
		ruleParts := strings.Split(line, ":")
		ruleId, _ := strconv.Atoi(ruleParts[0])

		if strings.Contains(ruleParts[1], "\"") {
			matchedValue := strings.Split(ruleParts[1], "\"")[1]
			rules[ruleId] = &ValueRule{value: matchedValue}
		} else {
			orParts := strings.Split(ruleParts[1], " | ")
			var ruleParts [][]int
			for _, part := range orParts {
				combinedIds, _ := utils.ToIntSlice(strings.Split(strings.Trim(part, " "), " "))
				ruleParts = append(ruleParts, combinedIds)
			}
			rules[ruleId] = &ComposedRule{rules: ruleParts}
		}
		lastRule = i
	}
	count := 0
	for _, line := range lines[lastRule+2:] {
		results := rules[0].match(line, &rules)
		if len(results) > 0 {
			count++
		}
	}
	return count, nil
}

func part2(lines []string) (int, error) {
	for i, line := range lines {
		if line == "8: 42" {
			lines[i] = "8: 42 | 42 8"
		} else if line == "11: 42 31" {
			lines[i] = "11: 42 31 | 42 11 31"
		}
	}

	return part1(lines)
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
