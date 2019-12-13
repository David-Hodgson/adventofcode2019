package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

func DayThirteenPartOne() {

	fmt.Println("2019 - Day Thirteen - Part One")

	input := strings.Split(ReadFile("day13-2019-input.txt"), ",")
	program := make([]int64, len(input)*10)

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = int64(value)
	}

	output := runProgram(program)

	blockCount := 0
	for i := 0; i < len(output); {

//		x := output[i]
		i++
//		y := output[i]
		i++
		tile_id := output[i]
		i++

		if tile_id == 2 {
			blockCount++
		}
	}
	fmt.Println("Block Count: ", blockCount)
}
