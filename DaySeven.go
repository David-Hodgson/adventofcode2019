package adventofcode2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DaySevenPartOne() {

	fmt.Println("2019 - Day Seven - Part One")

	input := strings.Split(ReadFile("day7-2019-input.txt"), ",")
	program := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = value
	}

	settings := generatePhaseSettings()

	inputSignal := 0
	outputValue := 0
	maxOutput := 0
	for s := 0; s < len(settings); s++ {
		inputSignal = 0
		phaseSetting := settings[s]

		for i :=0  ; i < 5; i++ {
			ampProgram := make([]int,len(program))
			copy(ampProgram,program)
			inputs := []int {phaseSetting[i],inputSignal}
			outputValue = d7RunProgram(ampProgram, inputs)
			inputSignal = outputValue
		}
		if outputValue > maxOutput {
			maxOutput = outputValue
		}
	}
	fmt.Println("Max Output: ", maxOutput)
}

func generatePhaseSettings() [][]int{

	settings := [][]int {}
	for i :=0; i < 5; i++ {

		for j :=0; j < 5; j++ {
			if i==j {
				continue
			}

			for k := 0; k < 5; k++ {
				if k==i || k==j {
					continue
				}

				for l :=0 ; l <  5; l++ {
					if l==i || l==j || l==k {
						continue
					}

					for m :=0; m<5; m++ {
						if m==i || m==j || m==k || m==l {
							continue
						}

						setting := []int {i,j,k,l,m}
						settings = append(settings,setting)
					}
				}
			}
		}

	}

	return settings
}

func d7RunProgram(program []int, inputs []int) int{

	halt := false

	output := -1
	for pc := 0; !halt; {

		instruction := program[pc]

		opcode := instruction % 100
		paramModes := instruction / 100
		switch opcode {
		case 1: //Add two numbers
			readOne := program[pc+1]
			pmOne := d7getParameterMode(paramModes, 0)
			readTwo := program[pc+2]
			pmTwo := d7getParameterMode(paramModes, 1)
			write := program[pc+3]

			program[write] = d7GetValue(readOne, pmOne, program) + getValue(readTwo, pmTwo, program)
			pc += 4
		case 2:
			readOne := program[pc+1]
			pmOne := d7getParameterMode(paramModes, 0)
			readTwo := program[pc+2]
			pmTwo := d7getParameterMode(paramModes, 1)
			write := program[pc+3]

			program[write] = d7GetValue(readOne, pmOne, program) * getValue(readTwo, pmTwo, program)
			pc += 4
		case 3:
			value := inputs[0]
			inputs = inputs[1:]
			param := program[pc+1]

			program[param] = value
			pc += 2
		case 4:
			param := program[pc+1]
			pmMode := d7getParameterMode(paramModes, 0)
			value := d7GetValue(param, pmMode, program)
			output = value
			pc += 2
		case 5:
			p1 := program[pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			v1 := d7GetValue(p1, pmMode1, program)
			v2 := d7GetValue(p2, pmMode2, program)

			if v1 != 0 {
				pc = v2
			} else {
				pc += 3
			}
		case 6:
			p1 := program[pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			v1 := d7GetValue(p1, pmMode1, program)
			v2 := d7GetValue(p2, pmMode2, program)

			if v1 == 0 {
				pc = v2
			} else {
				pc += 3
			}
		case 7:
			p1 := program[pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			p3 := program[pc+3]
			v1 := d7GetValue(p1, pmMode1, program)
			v2 := d7GetValue(p2, pmMode2, program)

			if v1 < v2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case 8:
			p1 := program[pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program[pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			p3 := program[pc+3]
			v1 := d7GetValue(p1, pmMode1, program)
			v2 := d7GetValue(p2, pmMode2, program)

			if v1 == v2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case 99:
			//fmt.Println("Halt")
			halt = true
		default:
			fmt.Println("Unknow Op Code: ", opcode)
			halt = true
		}

	}
	return output
}

func d7getParameterMode(paramModes int, param int) int {

	value := paramModes

	if param > 0 {
		value = value / int(math.Pow10(param))
	}
	return value % 10
}

func d7GetValue(index int, mode int, programMemory []int) int {

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
