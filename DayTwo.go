package adventofcode2019

import (
    "fmt"
    "strconv"
    "strings"
)

func DayTwoPartOne() {

    fmt.Println("2019 - Day Two - Part One")

    input := strings.Split(ReadFile("day2-2019-input.txt"),",")

    program := make([]int, len(input))

    for i := 0; i < len(input); i++ {
        value,_ := strconv.Atoi(input[i])
        program[i] = value
    }

    fmt.Println(program)

    program[1] = 12
    program[2] = 2

    fmt.Println("program[0]:", program[0])
    RunProgram(program)
    fmt.Println("program[0]:", program[0])
}

func DayTwoPartTwo() {

    fmt.Println("2019 - Day Two - Part Two")

    input := strings.Split(ReadFile("day2-2019-input.txt"),",")

    target := 19690720

    halt := false

    for i := 0 ; i < 100 && !halt; i++ {
        for j :=0; j < 100 && !halt; j++ {
            program := make([]int, len(input))


            for i := 0; i < len(input); i++ {

                value,_ := strconv.Atoi(input[i])
                program[i] = value

            }

            program[1] = i
            program[2] = j

            RunProgram(program)

            if program[0] == target {
                answer := 100 * i + j
                fmt.Println("Answer: ", answer)
                halt = true
            }
        }
    }
}

func RunProgram(program []int) {

    halt := false

    for pc := 0; !halt ; {

        opcode := program[pc]

        switch opcode {
        case 1:
            readOne:= program[pc+1]
            readTwo:= program[pc+2]
            write := program[pc+3]

            program[write] = program[readOne]+program[readTwo]
        case 2:
            readOne:= program[pc+1]
            readTwo:= program[pc+2]
            write := program[pc+3]

            program[write] = program[readOne] * program[readTwo]
        case 99:
            halt = true
        default:
            halt = true
        }

        pc += 4
    }
}
