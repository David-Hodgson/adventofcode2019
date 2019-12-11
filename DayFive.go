package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

func DayFivePartOne() {

	fmt.Println("2019 - Day Five - Part One")

	input := strings.Split(ReadFile("day5-2019-input.txt"), ",")
	program := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = value
	}

	fmt.Println(program)

	runProgram(program)
}


