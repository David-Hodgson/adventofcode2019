package adventofcode2019

import (
	"fmt"
	"strconv"
)

func DayEightPartOne() {

	fmt.Println("2019 - Day Eight - Part One")
	input := ReadFile("day8-2019-input.txt")

	width := 25
	height := 6

	layerCount := len(input) / (width * height)

	fmt.Println("Layer Count:", layerCount)

	lowestZeroCount := -1
	total := 0

	for layer := 0; layer < layerCount; layer++ {
		zeroCount := 0
		oneCount := 0
		twoCount := 0

		layerStart := layer * width * height
		layerEnd := layerStart + (width * height)
		layerValues := input[layerStart:layerEnd]

		for i := 0; i < len(layerValues); i++ {

			colour, _ := strconv.Atoi(string(layerValues[i]))

			switch colour {
			case 0:
				zeroCount++
			case 1:
				oneCount++
			case 2:
				twoCount++
			}
		}

		if lowestZeroCount == -1 || zeroCount < lowestZeroCount {
			lowestZeroCount = zeroCount
			total = oneCount * twoCount
		}
	}
	fmt.Println("Total: ", total)
}

func DayEightPartTwo() {

	fmt.Println("2019 - Day Eight - Part Two")

	input := ReadFile("day8-2019-input.txt")

	width := 25
	height := 6

	layerCount := len(input) / (width * height)

	fmt.Println("Layer Count:", layerCount)

	finalLayer := make([]int, height*width)

	for i := 0; i < len(finalLayer); i++ {
		finalLayer[i] = -1
	}

	printLayer(finalLayer, height, width)

	for layer := 0; layer < layerCount; layer++ {

		layerStart := layer * width * height
		layerEnd := layerStart + (width * height)
		layerValues := input[layerStart:layerEnd]

		for i := 0; i < len(layerValues); i++ {

			value, _ := strconv.Atoi(string(layerValues[i]))

			if finalLayer[i] == -1 && value != 2 {
				finalLayer[i] = value
			}
		}

		printLayer(finalLayer, height, width)
		fmt.Println("")
	}
}

func printLayer(layer []int, height int, width int) {

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {

			cell := layer[(row*width)+col]

			if cell != -1 && cell != 0 {
				fmt.Print(cell)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
