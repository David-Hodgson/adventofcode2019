package adventofcode2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ampController struct {
	pc           int
	instructions []int
	inputs       []int
}

func DaySevenPartOne() {

	fmt.Println("2019 - Day Seven - Part One")

	input := strings.Split(ReadFile("day7-2019-input.txt"), ",")
	program := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = value
	}

	settings := generatePhaseSettings(0, 5)

	inputSignal := 0
	outputValue := 0
	maxOutput := 0
	for s := 0; s < len(settings); s++ {
		inputSignal = 0
		phaseSetting := settings[s]

		for i := 0; i < 5; i++ {
			ampProgram := ampController{}
			ampProgram.pc = 0
			ampProgram.instructions = make([]int, len(program))

			ampProgram.inputs = []int{phaseSetting[i], inputSignal}
			copy(ampProgram.instructions, program)
			outputValue, _ = d7RunProgram(&ampProgram)
			inputSignal = outputValue
		}
		if outputValue > maxOutput {
			maxOutput = outputValue
		}
	}
	fmt.Println("Max Output: ", maxOutput)
}

func DaySevenPartTwo() {

	fmt.Println("2019 - Day Seven - Part Two")

	input := strings.Split(ReadFile("day7-2019-input.txt"), ",")
	program := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = value
	}

	settings := generatePhaseSettings(5, 10)

	inputSignal := 0
	outputValue := 0
	maxOutput := 0
	for s := 0; s < len(settings); s++ {
		inputSignal = 0
		phaseSetting := settings[s]

		amplifiers := make([]ampController, 5)
		for x := 0; x < 5; x++ {
			ampProgram := ampController{}
			ampProgram.pc = 0
			ampProgram.instructions = make([]int, len(program))
			ampProgram.inputs = []int{phaseSetting[x]}
			copy(ampProgram.instructions, program)
			amplifiers[x] = ampProgram
		}

		loops := 0
		for i := 0; loops < 100; {
			amplifiers[i].inputs = append(amplifiers[i].inputs, inputSignal)
			localoutputValue, halt := d7RunProgram(&amplifiers[i])

			inputSignal = localoutputValue
			outputValue = localoutputValue
			if i == 4 && halt {
				break
			}

			i++
			if i == 5 {
				i = 0
				loops++
			}
		}
		if outputValue > maxOutput {
			maxOutput = outputValue
		}
	}
	fmt.Println("Max Output: ", maxOutput)
}

func generatePhaseSettings(min int, max int) [][]int {

	settings := [][]int{}
	for i := min; i < max; i++ {

		for j := min; j < max; j++ {
			if i == j {
				continue
			}

			for k := min; k < max; k++ {
				if k == i || k == j {
					continue
				}

				for l := min; l < max; l++ {
					if l == i || l == j || l == k {
						continue
					}

					for m := min; m < max; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}

						setting := []int{i, j, k, l, m}
						settings = append(settings, setting)
					}
				}
			}
		}

	}

	return settings
}

func d7RunProgram(program *ampController) (int, bool) {

	halt := false
	wait := false
	output := -1
	for !halt && !wait {

		instruction := program.instructions[program.pc]

		opcode := instruction % 100
		paramModes := instruction / 100
		switch opcode {
		case 1: //Add two numbers
			readOne := program.instructions[program.pc+1]
			pmOne := d7getParameterMode(paramModes, 0)
			readTwo := program.instructions[program.pc+2]
			pmTwo := d7getParameterMode(paramModes, 1)
			write := program.instructions[program.pc+3]

			program.instructions[write] = d7GetValue(readOne, pmOne, program.instructions) + d7GetValue(readTwo, pmTwo, program.instructions)
			program.pc += 4
		case 2:
			readOne := program.instructions[program.pc+1]
			pmOne := d7getParameterMode(paramModes, 0)
			readTwo := program.instructions[program.pc+2]
			pmTwo := d7getParameterMode(paramModes, 1)
			write := program.instructions[program.pc+3]

			program.instructions[write] = d7GetValue(readOne, pmOne, program.instructions) * d7GetValue(readTwo, pmTwo, program.instructions)
			program.pc += 4
		case 3:

			if len(program.inputs) == 0 {
				wait = true
				break
			}
			value := program.inputs[0]
			program.inputs = program.inputs[1:]
			param := program.instructions[program.pc+1]

			program.instructions[param] = value
			program.pc += 2
		case 4:
			param := program.instructions[program.pc+1]
			pmMode := d7getParameterMode(paramModes, 0)
			value := d7GetValue(param, pmMode, program.instructions)
			output = value
			program.pc += 2
		case 5:
			p1 := program.instructions[program.pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program.instructions[program.pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			v1 := d7GetValue(p1, pmMode1, program.instructions)
			v2 := d7GetValue(p2, pmMode2, program.instructions)

			if v1 != 0 {
				program.pc = v2
			} else {
				program.pc += 3
			}
		case 6:
			p1 := program.instructions[program.pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program.instructions[program.pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			v1 := d7GetValue(p1, pmMode1, program.instructions)
			v2 := d7GetValue(p2, pmMode2, program.instructions)

			if v1 == 0 {
				program.pc = v2
			} else {
				program.pc += 3
			}
		case 7:
			p1 := program.instructions[program.pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program.instructions[program.pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			p3 := program.instructions[program.pc+3]
			v1 := d7GetValue(p1, pmMode1, program.instructions)
			v2 := d7GetValue(p2, pmMode2, program.instructions)

			if v1 < v2 {
				program.instructions[p3] = 1
			} else {
				program.instructions[p3] = 0
			}
			program.pc += 4
		case 8:
			p1 := program.instructions[program.pc+1]
			pmMode1 := d7getParameterMode(paramModes, 0)
			p2 := program.instructions[program.pc+2]
			pmMode2 := d7getParameterMode(paramModes, 1)
			p3 := program.instructions[program.pc+3]
			v1 := d7GetValue(p1, pmMode1, program.instructions)
			v2 := d7GetValue(p2, pmMode2, program.instructions)

			if v1 == v2 {
				program.instructions[p3] = 1
			} else {
				program.instructions[p3] = 0
			}
			program.pc += 4
		case 99:
			halt = true
		default:
			fmt.Println("Unknow Op Code: ", opcode)
			halt = true
		}

	}
	return output, halt
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
