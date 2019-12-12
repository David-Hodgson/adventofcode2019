package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

func DayNinePartOne() {

	fmt.Println("2019 - Day Nine - Part One")

	input := strings.Split(ReadFile("day9-2019-input.txt"), ",")
	program := make([]int64, len(input)*100)

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = int64(value)
	}

	runProgram(program)
}
