package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
	
)

type ingrediant struct {
	quanity int
	typeName string
}

type recipe struct {
	source []ingrediant
	output ingrediant
}

func DayFourteenPartOne() {
	fmt.Println("2019 - Day Fourteen - Part One")

	input := strings.Split(ReadFile("day14-2019-input.txt"), "\n")

	//fmt.Println("Input: ", input)

	recipes := make(map[string]recipe)

	for i := 0; i < len(input) ; i++ {
		recipe := parseRecipe(input[i])
		recipes[recipe.output.typeName] = recipe
	}

//	fmt.Println(recipes)

	target := "FUEL"

	r1 := recipes[target]
	fmt.Println("Quanity: ", r1.output.quanity)
	fmt.Println(r1.source)

	for i :=0; i<len(r1.source); i++ {
		fmt.Println("\t",r1.source[i].typeName)
		fmt.Println("\t\t", recipes[r1.source[i].typeName].output)
		fmt.Println("\t\t", recipes[r1.source[i].typeName].source)
	}
}

func parseRecipe(input string) recipe {
	recipeParts := strings.Split(input, "=>")

	sourceList := strings.Trim(recipeParts[0], " ")
	output := strings.Trim(recipeParts[1], " ")

	return recipe{parseIngrediantList(sourceList), parseIngrediant(output)}
}


func parseIngrediantList(input string) []ingrediant {

	ingrediantParts := strings.Split(input, ",")

	ingrediants := make([]ingrediant, len(ingrediantParts))

	for i := 0; i < len(ingrediantParts); i++ {
		ingrediants[i] = parseIngrediant(strings.Trim(ingrediantParts[i], " "))
	}

	return ingrediants
}

func parseIngrediant(input string) ingrediant {

	ingrediantParts := strings.Split(input, " ")

	quanity,_ := strconv.Atoi(ingrediantParts[0])
	typeName := ingrediantParts[1]

	return ingrediant{quanity, typeName}
}