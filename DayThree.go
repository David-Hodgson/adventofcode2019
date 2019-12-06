package adventofcode2019

import (
	"fmt"

	"strconv"

	"strings"
)

type line struct {
	start, end point
}

type point struct {
	x, y int
}

func DayThreePartOne() {

	fmt.Println("2019 - Day Three - Part One")

	//example1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"

	//example2 := "U62,R66,U55,R34,D71,R55,D58,R83"

	input := strings.Split(ReadFile("day3-2019-input.txt"), "\n")

	lines1 := parseLine(input[0])

	lines2 := parseLine(input[1])

	closestIntersection := -1

	for i := 0; i < len(lines1); i++ {

		for j := 0; j < len(lines2); j++ {

			intersect, x, y := calculateIntersection(lines1[i], lines2[j])

			if intersect && x != 0 && y != 0 {

				intersectDistance := getDistance(x, y)

				if closestIntersection == -1 ||

					intersectDistance < closestIntersection {

					closestIntersection = intersectDistance

				}

			}

		}

	}

	fmt.Println("Closest Intersection: ", closestIntersection)

}

func DayThreePartTwo() {

	fmt.Println("2019 - Day Three - Part Two")

	input := strings.Split(ReadFile("day3-2019-input.txt"), "\n")

	lines1 := parseLine(input[0])

	lines2 := parseLine(input[1])

	shortestLength := -1

	line1Length := 0

	for i := 0; i < len(lines1); i++ {

		line2Length := 0

		for j := 0; j < len(lines2); j++ {

			intersect, x, y := calculateIntersection(lines1[i], lines2[j])

			if intersect && x != 0 && y != 0 {

				line1ToIntsect := line{}

				line1ToIntsect.start = lines1[i].start

				line1ToIntsect.end = point{}

				line1ToIntsect.end.x = x

				line1ToIntsect.end.y = y

				line2ToIntsect := line{}

				line2ToIntsect.start = lines2[j].start

				line2ToIntsect.end = point{}

				line2ToIntsect.end.x = x

				line2ToIntsect.end.y = y

				l1Dist := getLineLength(line1ToIntsect)

				l2Dist := getLineLength(line2ToIntsect)

				currentLength := line1Length + line2Length + l1Dist + l2Dist

				if shortestLength == -1 ||

					currentLength < shortestLength {

					shortestLength = currentLength

				}

			}

			line2Length += getLineLength(lines2[j])

		}

		line1Length += getLineLength(lines1[i])

	}

	fmt.Println("Shortest Length: ", shortestLength)

}

func parseLine(rawInput string) []line {

	inputBits := strings.Split(rawInput, ",")

	x := 0

	y := 0

	lines := make([]line, len(inputBits))

	for i := 0; i < len(inputBits); i++ {

		dir := string(inputBits[i][0:1])

		distance, _ := strconv.Atoi(inputBits[i][1:])

		line := line{}

		line.start = point{x, y}

		switch dir {

		case "R":

			x = x + distance

		case "L":

			x = x - distance

		case "U":

			y = y + distance

		case "D":

			y = y - distance

		}

		line.end = point{x, y}

		lines[i] = line

	}

	return lines

}

func calculateIntersection(line1 line, line2 line) (bool, int, int) {

	//  fmt.Println("Checking lines")

	//  fmt.Println("\t", line1)

	//fmt.Println("\t", line2)

	if line1.start.x == line1.end.x {

		//line 1 is vertical

		if line2.start.x == line2.end.x {

			//line 2 is vertical

			//  fmt.Println("\t\tBoth Vertical")

			return false, 0, 0

		}

	}

	if line1.start.y == line1.end.y {

		//line 1 is horizontal

		if line2.start.y == line2.end.y {

			//line 2 is horizontal

			//  fmt.Println("\t\tBoth Horizontal")

			return false, 0, 0

		}

	}

	//One vertical line One horizontal line

	hLine := line{}

	vLine := line{}

	if line1.start.x == line1.end.x {

		//line 1 is vertical

		//fmt.Println("\t\tLine 1 Vertical")

		vLine = line1

		hLine = line2

	} else {

		//fmt.Println("\t\tLine 2 Vertical")

		vLine = line2

		hLine = line1

	}

	//intersecton point must be at

	// x: vline.start.x

	// y: hline.start.y

	x := vLine.start.x

	y := hLine.start.y

	//fmt.Println("\t\t\tIntersection point should be at x:",x," y:",y)

	vlineYmin := min(vLine.start.y, vLine.end.y)

	vlineYmax := max(vLine.start.y, vLine.end.y)

	hlineXMin := min(hLine.start.x, hLine.end.x)

	hlineXMax := max(hLine.start.x, hLine.end.x)

	if hlineXMin <= x && x <= hlineXMax &&

		vlineYmin <= y && y <= vlineYmax {

		return true, x, y

	}

	return false, 0, 0

}

func min(a, b int) int {

	if a <= b {

		return a

	}

	return b

}

func max(a, b int) int {

	if a >= b {

		return a

	}

	return b

}

func getDistance(x int, y int) int {

	a := x

	b := y

	if a < 0 {

		a = a * -1

	}

	if b < 0 {

		b = b * -1

	}

	return a + b

}

func getLineLength(line line) int {

	xMin := min(line.start.x, line.end.x)

	xMax := max(line.start.x, line.end.x)

	yMin := min(line.start.y, line.end.y)

	yMax := max(line.start.y, line.end.y)

	result := (xMax - xMin) + (yMax - yMin)

	return result

}
