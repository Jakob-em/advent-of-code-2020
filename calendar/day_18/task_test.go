package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestLexer(t *testing.T) {
	tokens := lex("(1 + 22) * 23 + ((1 + 1) * 2)")

	output := tokensToString(tokens)

	assert.Equal(t, "(1+22)*23+((1+1)*2)", output)
}

func TestPostfixWithoutWeights(t *testing.T) {
	tokens := lex("(1 + 22) * 23 + ((1 + 1) * 2)")
	postfix := convertToPostfix(tokens, func(i interface{}) int {
		switch i.(type) {
		case mult:
			return 1
		case add:
			return 1
		}
		return 0
	})
	output := tokensToString(postfix)

	assert.Equal(t, "122+23*11+2*+", output)
}

func TestPostfixWithWeights(t *testing.T) {
	tokens := []interface{}{
		par{open: true},
		number{val: 1},
		add{},
		number{val: 22},
		par{open: false},
		mult{},
		number{val: 23},
		add{},
		par{open: true},
		par{open: true},
		number{val: 1},
		add{},
		number{val: 1},
		par{open: false},
		mult{},
		number{val: 2},
		par{open: false},
	}
	postfix := convertToPostfix(tokens, func(i interface{}) int {
		switch i.(type) {
		case mult:
			return 1
		case add:
			return 2
		}
		return 0
	})
	output := tokensToString(postfix)

	assert.Equal(t, "122+2311+2*+*", output)
}

func TestEvaluatePostfix(t *testing.T) {
	assert.Equal(t, 24, evaluatePostfix([]interface{}{
		number{val: 5},
		number{val: 3},
		add{},
		number{val: 2},
		mult{},
		number{val: 3},
		number{val: 5},
		add{},
		add{},
	}))
}

func tokensToString(tokens []interface{}) string {
	output := ""
	for _, token := range tokens {
		switch t := token.(type) {
		case number:
			output += strconv.Itoa(t.val)
		case mult:
			output += "*"
		case add:
			output += "+"
		case par:
			if t.open {
				output += "("
			} else {
				output += ")"
			}
		}
	}
	return output
}
