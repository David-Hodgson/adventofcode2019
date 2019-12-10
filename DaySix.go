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

func DaySixPartTwo() {
	fmt.Println("2019 - Day Six - Part Two")

	input := strings.Split(ReadFile("day6-2019-input.txt"), "\n")
	tree := buildTree(input)

	depthMap := buildDepthMap(tree, "COM", 0)
	youDepth := depthMap["YOU"]
	sanDepth := depthMap["SAN"]
	pathCount := generatePathCount(tree,depthMap,"YOU","SAN")
	fmt.Println("Path Count: ", pathCount)
}

func generatePathCount(tree map[string][]string,depthMap map[string]int, node1 string, node2 string) int {

	node1Depth := depthMap[node1]
	node2Depth := depthMap[node2]

	if node1Depth == node2Depth {

		//nodes are at the same depth
		n1Parent := getNodeParent(node1,tree)
		n2Parent := getNodeParent(node2,tree)

		if n1Parent == n2Parent {
			//Same Parent
			return 0
		} else {
			return 2 + generatePathCount(tree,depthMap, n1Parent,n2Parent)
		}

	} else {

		//Mismatched nodes move the lowest one up
		if (node1Depth < node2Depth) {
			//move node 2
			n2Parent := getNodeParent(node2,tree)
			return 1 + generatePathCount(tree,depthMap, node1,n2Parent)
		} else {
			//move node 1
			n1Parent := getNodeParent(node1,tree)
			return 1 + generatePathCount(tree,depthMap, n1Parent,node2)
		}
	}

	return -1

}

func buildDepthMap(tree map[string][]string, root string, currentDepth int) map[string]int {

	depthMap := make(map[string]int)

	depthMap[root] = currentDepth

	children := tree[root]

	for i := 0; i < len(children); i++ {
	
		childMap := buildDepthMap(tree, children[i], currentDepth+1)

		for k,v := range childMap {
			depthMap[k] = v
		}
	}

	return depthMap
}

func getNodeParent(node string, tree map[string][]string) string {
	for k,v := range tree {

		for i := 0; i <len (v); i++ {
			if v[i] == node {
				return k
			}
		}

	}

	return "nil"
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
