package adventofcode2019

import (
	"fmt"
	"strings"
)

type asteroid struct {
	x, y int
}

type checkPoint struct {
	x, y int
}

func DayTenPartOne() {

	fmt.Println("2019 - Day Ten - Part One")

	input := strings.Split(ReadFile("day10-2019-input.txt"), "\n")
	asteroids := make(map[asteroid]bool, 0)

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {

			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}

			value := string(input[y][x])

			if value == "#" {
				asteroid := asteroid{x, y}
				asteroids[asteroid] = true
			}
		}
	}

	highestCount := 0

	for a, _ := range asteroids {

		count := generateCountForAsteroid(a, asteroids, minX, maxX, minY, maxY)
		if count > highestCount {
			highestCount = count
		}
	}
	fmt.Println("Count: ", highestCount)
}

func generateCountForAsteroid(asteroid asteroid, starMap map[asteroid]bool, minX, maxX, minY, maxY int) int {

	asteroidCount := 0
	checkMap := make(map[checkPoint]bool)
	asteroidCount += checkBottomRightQuad(asteroid, starMap, checkMap, minX, maxX, minY, maxY)
	asteroidCount += checkBottomLeftQuad(asteroid, starMap, checkMap, minX, maxX, minY, maxY)
	asteroidCount += checkTopRightQuad(asteroid, starMap, checkMap, minX, maxX, minY, maxY)
	asteroidCount += checkTopLeftQuad(asteroid, starMap, checkMap, minX, maxX, minY, maxY)
	return asteroidCount
}

func checkBottomRightQuad(asteroid asteroid, starMap map[asteroid]bool, checkMap map[checkPoint]bool, minX, maxX, minY, maxY int) int {

	asteroidCount := 0
	for deltaX := 0; deltaX <= (maxX - asteroid.x); deltaX++ {
		for deltaY := 0; deltaY <= (maxY - asteroid.y); deltaY++ {
			if deltaX == 0 && deltaY == 0 {
				continue
			}

			asteroidCount += checkLine(asteroid, starMap, checkMap, deltaX, deltaY, minX, maxX, minY, maxY)
		}
	}
	return asteroidCount
}

func checkBottomLeftQuad(asteroid asteroid, starMap map[asteroid]bool, checkMap map[checkPoint]bool, minX, maxX, minY, maxY int) int {

	asteroidCount := 0
	for deltaX := 0; deltaX >= (minX - asteroid.x); deltaX-- {
		for deltaY := 0; deltaY <= (maxY - asteroid.y); deltaY++ {
			if deltaX == 0 && deltaY == 0 {
				continue
			}

			asteroidCount += checkLine(asteroid, starMap, checkMap, deltaX, deltaY, minX, maxX, minY, maxY)
		}
	}
	return asteroidCount
}

func checkTopRightQuad(asteroid asteroid, starMap map[asteroid]bool, checkMap map[checkPoint]bool, minX, maxX, minY, maxY int) int {

	asteroidCount := 0
	for deltaX := 0; deltaX <= (maxX - asteroid.x); deltaX++ {
		for deltaY := 0; deltaY >= (minY - asteroid.y); deltaY-- {
			if deltaX == 0 && deltaY == 0 {
				continue
			}

			asteroidCount += checkLine(asteroid, starMap, checkMap, deltaX, deltaY, minX, maxX, minY, maxY)
		}
	}
	return asteroidCount
}

func checkTopLeftQuad(asteroid asteroid, starMap map[asteroid]bool, checkMap map[checkPoint]bool, minX, maxX, minY, maxY int) int {

	asteroidCount := 0
	for deltaX := 0; deltaX >= (minX - asteroid.x); deltaX-- {
		for deltaY := 0; deltaY >= (minY - asteroid.y); deltaY-- {
			if deltaX == 0 && deltaY == 0 {
				continue
			}

			asteroidCount += checkLine(asteroid, starMap, checkMap, deltaX, deltaY, minX, maxX, minY, maxY)
		}
	}
	return asteroidCount
}

func checkLine(currentAsteroid asteroid, starMap map[asteroid]bool, checkMap map[checkPoint]bool, deltaX, deltaY, minX, maxX, minY, maxY int) int {

	x := currentAsteroid.x // + deltaX;
	y := currentAsteroid.y // + deltaY;
	found := false
	for x <= maxX && y <= maxY && y >= minY && x >= minX {

		x = x + deltaX
		y = y + deltaY
		checkPoint := checkPoint{x, y}
		if _, exists := checkMap[checkPoint]; exists {
			break
		}

		checkMap[checkPoint] = true

		checkAsteroid := asteroid{x, y}
		if !found && starMap[checkAsteroid] == true {
			found = true
		}
	}

	if found {
		return 1
	} else {
		return 0
	}
}
