package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
	
)

type vector struct {
	x,y,z int
}

type moon struct {
	position,velocity vector
}

func DayTwelvePartOne() {
	fmt.Println("2019 - Day Twelve - Part One")

	input := strings.Split(ReadFile("day12-2019-input.txt"), "\n")
	
	moons := make([]moon, len(input))
	for i :=0; i < len(input); i++ {

		moons[i] = parseMoonInput(input[i])
	}

	for t := 0; t < 1000; t++ {
		updateMoons(moons)
	}

	fmt.Println("Total Energy: ", getEnergy(moons))
}

func getEnergy(moons []moon) int {
	totalEnergy := 0
	for m := 0; m <len(moons); m++ {

		moon := moons[m]
		potential := abs(moon.position.x) + abs(moon.position.y) + abs(moon.position.z)
		kinetic := abs(moon.velocity.x) + abs(moon.velocity.y) + abs(moon.velocity.z)

		totalEnergy = totalEnergy + (potential * kinetic)

	}
	return totalEnergy;
}

func updateMoons(moons []moon) {

	applyGravity(moons)
	applyVelocity(moons)
}

func applyGravity(moons []moon) {

	for m1 := 0; m1 <len(moons); m1++ {
		for m2:=m1+1; m2<len(moons); m2++ {

			moon1 := &moons[m1]
			moon2 := &moons[m2]

			applyXGravity(moon1, moon2)
			applyYGravity(moon1, moon2)
			applyZGravity(moon1, moon2)
		}
	}
}

func applyXGravity(moon1, moon2 *moon) {
	if (moon1.position.x < moon2.position.x) {
		moon1.velocity.x = moon1.velocity.x + 1
		moon2.velocity.x = moon2.velocity.x - 1
	}

	if (moon1.position.x > moon2.position.x) {
		moon1.velocity.x = moon1.velocity.x - 1
		moon2.velocity.x = moon2.velocity.x + 1
	}
}

func applyYGravity(moon1, moon2 *moon) {
	if (moon1.position.y < moon2.position.y) {
		moon1.velocity.y = moon1.velocity.y + 1
		moon2.velocity.y = moon2.velocity.y - 1
	}

	if (moon1.position.y > moon2.position.y) {
		moon1.velocity.y = moon1.velocity.y - 1
		moon2.velocity.y = moon2.velocity.y + 1
	}
}

func applyZGravity(moon1, moon2 *moon) {
	if (moon1.position.z < moon2.position.z) {
		moon1.velocity.z = moon1.velocity.z + 1
		moon2.velocity.z = moon2.velocity.z - 1
	}

	if (moon1.position.z> moon2.position.z) {
		moon1.velocity.z = moon1.velocity.z - 1
		moon2.velocity.z = moon2.velocity.z + 1
	}
}

func applyVelocity(moons []moon) {
	for m := 0; m < len(moons); m++ {
		moon := &moons[m]
		moon.position.x = moon.position.x + moon.velocity.x
		moon.position.y = moon.position.y + moon.velocity.y
		moon.position.z = moon.position.z + moon.velocity.z
		
	}
}

func parseMoonInput(input string) moon{
	dataString := input

	if dataString[0:1] == "<" {
		dataString = dataString[1:]
	}
	if dataString[len(dataString)-1:len(dataString)] == ">" {
		dataString = dataString[0: len(dataString)-1]
	}

	vectorParts := strings.Split(dataString, ",")

	x := 0
	y := 0
	z := 0

	for i :=0; i < len(vectorParts); i++ {
		parts := strings.Split(strings.Trim(vectorParts[i], " "), "=")
		coord := string(parts[0])
		value,_ := strconv.Atoi(parts[1])

		switch (coord) {
		case "x":
			x = value
		case "y":
			y = value
		case "z":
			z = value
	
		}
		
	}

	position := vector{x,y,z}
	moon := moon{}
	moon.position = position

	return moon
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}