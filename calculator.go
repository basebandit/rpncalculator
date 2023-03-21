package rpncalculator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/basebandit/rpncalculator/util"
)

// Evaluate checks the space delimited string expression
func Evaluate(expr string) (float64, error) {
	if len(expr) == 0 {
		return 0, nil
	}

	var start, end int // 0
	stack := util.NewStack()
	for start < len(expr) {
		// find the next space
		space := strings.Index(expr[start:], " ")
		if space == -1 {
			end = len(expr)
		} else {
			end = start + space
		}

		current := expr[start:end]

		// if current is an operator, apply the operation to top 2 operands on the stack.
		if strings.IndexAny(current, "+-*/") != -1 {
			opA, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("evaluate error: %w", err)
			}
			opB, err := stack.Pop()
			if err != nil {
				return 0, fmt.Errorf("evaluate error: %w", err)
			}
			stack.Push(operate(current[0], opB.(float64), opA.(float64)))
		} else { //otherwise push it to stack
			num, err := strconv.ParseFloat(current, 64)
			if err != nil {
				return 0, fmt.Errorf("evaluate error: %w. could be the operator is not supported yet", err)
			}
			stack.Push(num)
		}
		start = end + 1 // start over at index after the space.
	}

	result, err := stack.Pop()
	if err != nil {
		return 0, fmt.Errorf("evaluate error: %w", err)
	}

	for !stack.Empty() {
		current, err := stack.Pop()
		if err != nil {
			return 0, fmt.Errorf("evaluate error: %w", err)
		}
		if current.(float64) > result.(float64) {
			result = current
		}
	}
	return result.(float64), nil
}

// operate returns a <operator> b
func operate(operator uint8, operandA, operandB float64) float64 {
	opTable := make(map[rune]float64)
	opTable['+'] = operandA + operandB
	opTable['-'] = operandA - operandB
	opTable['*'] = operandA * operandB
	opTable['/'] = operandA / operandB
	return opTable[rune(operator)]
}
