package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

func DayOnePartOne() {
	fmt.Println("2019 - Day One - Part One")

	input := strings.Split(ReadFile("day1-2019-input.txt"),"\n")

	fuelRequired := 0

	for i := 0; i<len(input); i++ {
		module, _ := strconv.Atoi(string(input[i]))

		fmt.Println("Module: ", module)

		moduleFuel := (module / 3)-2

		fmt.Println("Module Fuel: ", moduleFuel)

		fuelRequired += moduleFuel
	}

	fmt.Println("Total Fuel: ", fuelRequired)
}


func DayOnePartTwo() {
	fmt.Println("2019 - Day One - Part Two")

	input := strings.Split(ReadFile("day1-2019-input.txt"),"\n")

	fuelRequired := 0

	for i := 0; i<len(input); i++ {
		module, _ := strconv.Atoi(string(input[i]))

		fmt.Println("Module: ", module)

		moduleFuel := (module / 3)-2

		fmt.Println("Module Fuel: ", moduleFuel)

		moduleFuel += getFuel(moduleFuel)
		fuelRequired += moduleFuel 
	}

	fmt.Println("Total Fuel: ", fuelRequired)
}

func getFuel(wieght int) int {

	fuel := (wieght /3) - 2

	if fuel > 0 {
		fuel += getFuel(fuel)
	} else {
		fuel = 0
	}

	return fuel
}
