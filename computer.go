package adventofcode2019

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func runProgramViaCallBack(program []int64, inputCallback func()int64, outputCallback func(int64)) []int64 {

	fmt.Println("Running")
	halt := false
	relativeBase := int64(0)
	output := []int64{}

	for pc := int64(0); !halt; {

		instruction := program[pc]

		opcode := instruction % 100
		paramModes := instruction / 100

		switch opcode {
		case 1: //Add two numbers
			readOne := program[pc+1]
			pmOne := getParameterMode(paramModes, 0)
			readTwo := program[pc+2]
			pmTwo := getParameterMode(paramModes, 1)

			pmThree := getParameterMode(paramModes, 2)

			write := program[pc+3]

			value := getValue(readOne, pmOne, relativeBase, program) + getValue(readTwo, pmTwo, relativeBase, program)
			setValue(value, write, pmThree, relativeBase, program)
			pc += 4
		case 2:
			readOne := program[pc+1]
			pmOne := getParameterMode(paramModes, 0)
			readTwo := program[pc+2]
			pmTwo := getParameterMode(paramModes, 1)
			pmThree := getParameterMode(paramModes, 2)
			write := program[pc+3]

			value := getValue(readOne, pmOne, relativeBase, program) * getValue(readTwo, pmTwo, relativeBase, program)
			setValue(value, write, pmThree, relativeBase, program)
			pc += 4
		case 3:
			value := inputCallback()

			param := program[pc+1]
			paramMode := getParameterMode(paramModes, 0)
			setValue(value, param, paramMode, relativeBase, program)
			pc += 2
		case 4:
			param := program[pc+1]
			pmMode := getParameterMode(paramModes, 0)
			value := getValue(param, pmMode, relativeBase, program)

			outputCallback(value)
			pc += 2
		case 5:
			p1 := program[pc+1]
			pmMode1 := getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := getParameterMode(paramModes, 1)
			v1 := getValue(p1, pmMode1, relativeBase, program)
			v2 := getValue(p2, pmMode2, relativeBase, program)

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
			v1 := getValue(p1, pmMode1, relativeBase, program)
			v2 := getValue(p2, pmMode2, relativeBase, program)

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

			pmMode3 := getParameterMode(paramModes, 2)

			v1 := getValue(p1, pmMode1, relativeBase, program)
			v2 := getValue(p2, pmMode2, relativeBase, program)

			if v1 < v2 {
				setValue(1, p3, pmMode3, relativeBase, program)
			} else {
				setValue(0, p3, pmMode3, relativeBase, program)
			}
			pc += 4
		case 8:
			p1 := program[pc+1]
			pmMode1 := getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := getParameterMode(paramModes, 1)
			p3 := program[pc+3]

			pmThree := getParameterMode(paramModes, 2)
			v1 := getValue(p1, pmMode1, relativeBase, program)
			v2 := getValue(p2, pmMode2, relativeBase, program)

			if v1 == v2 {
				setValue(1, p3, pmThree, relativeBase, program)
			} else {
				setValue(0, p3, pmThree, relativeBase, program)
			}
			pc += 4
		case 9:
			p1 := program[pc+1]
			pmMode1 := getParameterMode(paramModes, 0)
			v1 := getValue(p1, pmMode1, relativeBase, program)
			relativeBase = relativeBase + v1

			pc += 2
		case 99:
			fmt.Println("Halt")
			halt = true
		default:
			fmt.Println("Unknow Op Code: ", opcode)
			halt = true
		}

	}
	return output
}


func runProgram(program []int64) []int64 {

	inputCallback := func() int64 {
		fmt.Println("Input:")
		reader := bufio.NewReader(os.Stdin)
		char, _, _ := reader.ReadRune()
		value, _ := strconv.Atoi(string(char))

		return int64(value)
	}

	
	output := []int64{}
	outputCallback := func(outputValue int64) {
		if output == nil {
			fmt.Println("Output: ", outputValue)
		} else {
			output = append(output,outputValue)
		}

	}

	runProgramViaCallBack(program,inputCallback,outputCallback)
	return output
}

func getParameterMode(paramModes int64, param int64) int64 {

	value := paramModes

	if param > 0 {
		value = value / int64(math.Pow10(int(param)))
	}
	return value % 10
}

func getValue(index int64, mode int64, relativeBase int64, programMemory []int64) int64 {

	if mode == 0 {
		//position mode
		return programMemory[index]
	}

	if mode == 1 {
		//immediate mode
		return index
	}

	if mode == 2 {
		//relative mode
		return programMemory[index+relativeBase]
	}
	return -1
}

func setValue(value int64, index int64, mode int64, relativeBase int64, programMemory []int64) {
	if mode == 0 {
		//position mode
		programMemory[index] = value
	}

	if mode == 2 {
		programMemory[index+relativeBase] = value
	}
}
