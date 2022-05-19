package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var plusIndex, minusIndex, result, plus, minus int

	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return "", fmt.Errorf("input error: %w", errorEmptyInput)
	}

	plus = strings.Count(input, "+")
	plusIndex = strings.Index(input, "+")
	minus = strings.Count(input, "-")
	minusIndex = strings.Index(input, "-")

	if plus > 1 || minus > 2 || (plus+minus) > 2 || (plus+minus) == 0 || (plus+minus) > 1 && minusIndex != 0 {
		return "", fmt.Errorf("expression error: %w", errorNotTwoOperands)
	}

	if plus == 1 {
		if len(input[:plusIndex]) == 0 || len(input[plusIndex+1:]) == 0 {
			return "", fmt.Errorf("expression error: %w", errorNotTwoOperands)
		}
		first, err1 := strconv.Atoi(input[:plusIndex])
		if err1 != nil {
			return "", fmt.Errorf("first operand error: %w", err1)
		}
		second, err2 := strconv.Atoi(input[plusIndex+1:])
		if err2 != nil {
			return "", fmt.Errorf("second operand error: %w", err2)
		}
		result = first + second
	} else {
		if minusIndex == 0 {
			minusIndex = strings.LastIndex(input, "-")
		}
		if len(input[:minusIndex]) == 0 || len(input[minusIndex+1:]) == 0 {
			return "", fmt.Errorf("expression error: %w", errorNotTwoOperands)
		}
		first, err1 := strconv.Atoi(input[:minusIndex])
		if err1 != nil {
			return "", fmt.Errorf("first operand error: %w", err1)
		}
		second, err2 := strconv.Atoi(input[minusIndex+1:])
		if err2 != nil {
			return "", fmt.Errorf("second operand error: %w", err2)
		}
		result = first - second
	}

	return strconv.Itoa(result), nil
}
