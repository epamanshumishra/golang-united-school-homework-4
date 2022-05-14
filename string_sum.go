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
	var sum int
	var operator rune
	var intCount int
	var cnt int
	var tmpstr string
	stringWithoutSpaces := strings.ReplaceAll(input, " ", "")
	if stringWithoutSpaces == "" {
		errorEmptyString := fmt.Errorf("empty input:%w", errorEmptyInput)
		return "", errorEmptyString
	}
	for pos := 0; pos < len(stringWithoutSpaces); pos++ {
		if input[pos] == '-' {
			operator = '-'
		} else if input[pos] == '+' {
			operator = '+'
		} else if stringWithoutSpaces[pos] >= '0' && stringWithoutSpaces[pos] <= '9' {
			for i := pos; i < len(stringWithoutSpaces) && stringWithoutSpaces[i] >= '0' && stringWithoutSpaces[i] <= '9'; i++ {
				cnt++
			}
			tmpstr = stringWithoutSpaces[pos : pos+cnt]
			intVar, err := strconv.Atoi(tmpstr)
			if err != nil {
				return "", fmt.Errorf("input expression is not valid(contains characters, that are not numbers, +, - or whitespace): %w", err)
			}
			if operator == '+' {
				sum += intVar
			} else if operator == '-' {
				sum -= intVar
			} else {
				sum = intVar
			}
			pos += cnt
			cnt = 0
			intCount++
		} else {
			return "", fmt.Errorf("input expression is not valid(contains characters, that are not numbers, +, - or whitespace)")
		}
	}
	if intCount != 2 {
		errorCountOperands := fmt.Errorf("less or more than two operands: %w", errorNotTwoOperands)
		return "", errorCountOperands
	}
	return strconv.Itoa(sum), nil
}
