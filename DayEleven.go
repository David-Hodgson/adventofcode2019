package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	UP =iota
	DOWN = iota
	LEFT = iota
	RIGHT = iota
)

type paintingRobot struct {
	x,y,currentDirection int
}

type point2d struct {
	x,y int
}

var lock = sync.RWMutex{}

func DayElvenPartOne() {
	fmt.Println("2019 - Day Eleven - Part One")

	input := strings.Split(ReadFile("day11-2019-input.txt"), ",")
	program := make([]int64, len(input)*1000)

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = int64(value)
	}

	robot := paintingRobot{10000,10000, UP}
	hull := make(map[point2d]int64)

	inputCallback := func() int64 {
		robotLocation := point2d{robot.x,robot.y}
		paintValue := hull[robotLocation]

		return paintValue
	}

	paintMode := true
	outputCallback := func(outputValue int64) {
			if paintMode {
				robotLocation := point2d{robot.x,robot.y}
				hull[robotLocation] = int64(outputValue)
			} else {
				newDirection := -1
				if outputValue ==0 {
					newDirection = turnLeft(robot.currentDirection)
				} else  {
					newDirection = turnRight(robot.currentDirection)
				}
	
				robot.currentDirection = newDirection
				
				moveRobot(&robot)
			}
			paintMode = !paintMode
	}

	runProgramViaCallBack(program,inputCallback, outputCallback)

	paintCount := len(hull);
	fmt.Println("Paint Count:", paintCount)
	
}


func DayElvenPartTwo() {
	fmt.Println("2019 - Day Eleven - Part One")

	input := strings.Split(ReadFile("day11-2019-input.txt"), ",")
	program := make([]int64, len(input)*1000)

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		program[i] = int64(value)
	}

	robot := paintingRobot{0,0, UP}
	hull := make(map[point2d]int64)
	hull[point2d{0,0}] = 1

	inputCallback := func() int64 {
		robotLocation := point2d{robot.x,robot.y}
		paintValue := hull[robotLocation]

		return paintValue
	}

	maxX := robot.x
	minX := robot.x
	maxY := robot.y
	minY := robot.y

	paintMode := true
	outputCallback := func(outputValue int64) {
			if paintMode {
				robotLocation := point2d{robot.x,robot.y}
				hull[robotLocation] = int64(outputValue)
			} else {
				newDirection := -1
				if outputValue ==0 {
					newDirection = turnLeft(robot.currentDirection)
				} else  {
					newDirection = turnRight(robot.currentDirection)
				}
	
				robot.currentDirection = newDirection
				
				moveRobot(&robot)
				if robot.x <minX {
					minX = robot.x
				}
				if robot.x > maxX {
					maxX = robot.x
				}
				if robot.y <minY {
					minY = robot.y
				}
				if robot.y > maxY{
					maxY = robot.y
				}


			}
			paintMode = !paintMode
	}

	runProgramViaCallBack(program,inputCallback, outputCallback)

	paintCount := len(hull);
	fmt.Println("Paint Count:", paintCount)
	
	fmt.Println("maxX: ", maxX)
	fmt.Println("minX: ", minX)
	fmt.Println("maxY: ", maxY)
	fmt.Println("minY: ", minY)

	for y :=maxY; y >= minY; y-- {

		for x := minX; x<maxX;x++ {
			value := hull[point2d{x,y}]
			if value == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func moveRobot(robot *paintingRobot) {
	
	switch(robot.currentDirection) {
	case UP:
		robot.y = robot.y + 1
	case DOWN:
		robot.y = robot.y - 1
	case LEFT:
		robot.x = robot.x - 1
	case RIGHT:
		robot.x = robot.x + 1
	}
	
}

func turnLeft(currentDirection int) int {

	switch (currentDirection) {
	case UP:
		return LEFT
	case DOWN:
		return RIGHT
	case LEFT:
		return DOWN
	case RIGHT:
		return UP
	}

	return -1
}

func turnRight(currentDirection int) int {

	switch (currentDirection) {
	case UP:
		return RIGHT
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	}

	return -1
}


