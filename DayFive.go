package adventofcode2019

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func runProgram(program []int) {

	halt := false

	for pc := 0; !halt; {

		instruction := program[pc]

		opcode := instruction % 100
		paramModes := instruction / 100
		switch opcode {
		case 1: //Add two numbers
			readOne := program[pc+1]
			pmOne := getParameterMode(paramModes, 0)
			readTwo := program[pc+2]
			pmTwo := getParameterMode(paramModes, 1)
			write := program[pc+3]

			program[write] = getValue(readOne, pmOne, program) + getValue(readTwo, pmTwo, program)
			pc += 4
		case 2:
			readOne := program[pc+1]
			pmOne := getParameterMode(paramModes, 0)
			readTwo := program[pc+2]
			pmTwo := getParameterMode(paramModes, 1)
			write := program[pc+3]

			program[write] = getValue(readOne, pmOne, program) * getValue(readTwo, pmTwo, program)
			pc += 4
		case 3:
			fmt.Println("Input:")
			reader := bufio.NewReader(os.Stdin)
			char, _, _ := reader.ReadRune()
			value, _ := strconv.Atoi(string(char))
			param := program[pc+1]

			program[param] = value
			pc += 2
		case 4:
			param := program[pc+1]
			pmMode := getParameterMode(paramModes, 0)
			value := getValue(param, pmMode, program)
			fmt.Println("Output: ", value)
			pc += 2
		case 5:
			p1 := program[pc+1]
			pmMode1 := getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := getParameterMode(paramModes, 1)
			v1 := getValue(p1, pmMode1, program)
			v2 := getValue(p2, pmMode2, program)

			if v1 != 0 {
				pc = v2
			} else {
				pc += 3
			}
		case 6:
			p1 := program[pc+1]
			pmMode1 := getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := getParameterMode(paramModes, 1)
			v1 := getValue(p1, pmMode1, program)
			v2 := getValue(p2, pmMode2, program)

			if v1 == 0 {
				pc = v2
			} else {
				pc += 3
			}
		case 7:
			p1 := program[pc+1]
			pmMode1 := getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := getParameterMode(paramModes, 1)
			p3 := program[pc+3]
			v1 := getValue(p1, pmMode1, program)
			v2 := getValue(p2, pmMode2, program)

			if v1 < v2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case 8:
			p1 := program[pc+1]
			pmMode1 := getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := getParameterMode(paramModes, 1)
			p3 := program[pc+3]
			v1 := getValue(p1, pmMode1, program)
			v2 := getValue(p2, pmMode2, program)

			if v1 == v2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case 99:
			fmt.Println("Halt")
			halt = true
		default:
			fmt.Println("Unknow Op Code: ", opcode)
			halt = true
		}

	}
}

func getParameterMode(paramModes int, param int) int {

	value := paramModes

	if param > 0 {
		value = value / int(math.Pow10(param))
	}
	return value % 10
}

func getValue(index int, mode int, programMemory []int) int {

	if mode == 0 {
		//position mode
		return programMemory[index]
	}

	if mode == 1 {
		//immediate mode
		return index
	}

	return -1
}
