package adventofcode2019

import (
	"fmt"
	"strings"
)

func DaySixPartOne() {
	fmt.Println("2019 - Day Six - Part One")

	input := strings.Split(ReadFile("day6-2019-input.txt"), "\n")
	tree := buildTree(input)
	total := calculateOrbits(tree, "COM", 0)
	fmt.Println("total: ", total)
}

func calculateOrbits(tree map[string][]string, start string, depth int) int {

	total := 0
	childNodes := tree[start]

	for i := 0; i < len(childNodes); i++ {
		total += depth + 1
		total += calculateOrbits(tree, childNodes[i], depth+1)
	}

	return total
}

func buildTree(input []string) map[string][]string {

	tree := make(map[string][]string)

	for i := 0; i < len(input); i++ {

		bits := strings.Split(input[i], ")")

		_, exists := tree[bits[0]]

		if !exists {
			tree[bits[0]] = []string{}
		}
		tree[bits[0]] = append(tree[bits[0]], bits[1])
	}

	return tree
}
