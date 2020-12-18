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

func part1(lines []string) (int, error) {

	sum := 0
	for _, line := range lines {
		sum += evaluatePostfix(convertToPostfix(lex(line), func(i interface{}) int {
			switch i.(type) {
			case mult:
				return 1
			case add:
				return 1
			}
			return 0
		}))
	}

	return sum, nil
}

func part2(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		sum += evaluatePostfix(convertToPostfix(lex(line), func(i interface{}) int {
			switch i.(type) {
			case mult:
				return 1
			case add:
				return 2
			}
			return 0
		}))
	}

	return sum, nil
}

type number struct {
	val int
}

type par struct {
	open bool
}

type add struct {
}

type mult struct {
}

func lex(input string) []interface{} {
	var tokens []interface{}
	currentToken := ""
	for _, c := range strings.Replace(input, " ", "", -1) {
		switch c {
		case '(':
			tokens = append(tokens, par{open: true})
		case ')':
			tokens = addNumber(&currentToken, tokens)
			tokens = append(tokens, par{open: false})
		case '+':
			tokens = addNumber(&currentToken, tokens)
			tokens = append(tokens, add{})
		case '*':
			tokens = addNumber(&currentToken, tokens)
			tokens = append(tokens, mult{})
		default:
			currentToken += string(c)
		}
	}
	tokens = addNumber(&currentToken, tokens)
	return tokens
}

func addNumber(currentToken *string, tokens []interface{}) []interface{} {
	num, err := strconv.Atoi(*currentToken)
	if err != nil {
		return tokens
	}
	*currentToken = ""
	tokens = append(tokens, number{val: num})
	return tokens
}

type precedenceMapper func(interface{}) int

//convertToPostfix Convert token list from infix to postfix notation using the shunting yard algorithm
func convertToPostfix(tokens []interface{}, precedence precedenceMapper) []interface{} {
	var output []interface{}
	var operators utils.Stack
	for _, token := range tokens {
		switch t := token.(type) {
		case number:
			output = append(output, t)
		case par:
			if t.open {
				operators.Push(t)
			} else {
				for operators.Peek() != (par{open: true}) {
					output = append(output, operators.Pop())
				}
				operators.Pop()
			}
		case mult, add:
			for tos := operators.Peek(); isOperator(tos) && precedence(tos) >= precedence(t); tos = operators.Peek() {
				output = append(output, operators.Pop())
			}
			operators.Push(t)
		}
	}
	for operators.Peek() != nil {
		output = append(output, operators.Pop())
	}
	return output
}

func evaluatePostfix(postfix []interface{}) int {
	var s utils.Stack
	for _, token := range postfix {
		switch t := token.(type) {
		case number:
			s.Push(t)
		case mult:
			s.Push(number{val: s.Pop().(number).val * s.Pop().(number).val})
		case add:
			s.Push(number{val: s.Pop().(number).val + s.Pop().(number).val})
		}
	}
	return s.Pop().(number).val
}

func isOperator(val interface{}) bool {
	switch val.(type) {
	case add:
		return true
	case mult:
		return true
	default:
		return false
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
