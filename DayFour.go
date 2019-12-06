package adventofcode2019

import (
	"fmt"
	"strconv"
	//"strings"
)

/*
However, they do remember a few key facts about the password:
It is a six-digit number.
The value is within the range given in your puzzle input.
Two adjacent digits are the same (like 22 in 122345).
Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
Other than the range rule, the following are true:
111111 meets these criteria (double 11, never decreases).
223450 does not meet these criteria (decreasing pair of digits 50).
123789 does not meet these criteria (no double).

How many different passwords within the range given in your puzzle input meet these criteria?

*/

func DayFourPartOne() {
	fmt.Println("2019 - Day Four - Part One")

	example1 := "111111"
	example2 := "223450"
	example3 := "123789"

	fmt.Println("Example 1 isValid: ", isValid(example1))
	fmt.Println("Example 2 isValid: ", isValid(example2))
	fmt.Println("Example 3 isValid: ", isValid(example3))

	minValue := 168630
	maxValue := 718098

	validCount := 0

	for i := minValue; i <= maxValue; i++ {
		password := strconv.Itoa(i)

		if isValid(password) {
			validCount++
		}
	}

	fmt.Println("Valid Password Count: ", validCount)
}

func isValid(password string) bool {

	isValid := true

	if !isSixDigits(password) {
		isValid = false
	}

	if !hasTwoSameCharacters(password) {
		isValid = false
	}

	if !hasNoDescreasingValues(password) {
		isValid = false
	}

	return isValid
}

func isSixDigits(password string) bool {

	if len(password) != 6 {
		return false
	}
	return true
}

func isWithinRange(password string, min, max int) bool {

	value, _ := strconv.Atoi(password)

	if value >= min && value <= max {
		return true
	}

	return false
}

func hasTwoSameCharacters(password string) bool {

	valid := false

	for i := 0; i < len(password)-1; i++ {

		if password[i] == password[i+1] {
			valid = true
		}
	}
	return valid
}

func hasNoDescreasingValues(password string) bool {

	valid := true

	for i := 0; i < len(password)-1; i++ {
		value1, _ := strconv.Atoi(string(password[i]))
		value2, _ := strconv.Atoi(string(password[i+1]))

		if value2 < value1 {
			valid = false
		}
	}
	return valid
}
