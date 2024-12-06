package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ruleDictionary, updates := parseInput("input.txt")

	sumOfMiddlePages := 0
	var invalidUpdates []Update

	for _, update := range updates {
		isValid := true

		for x, ruleMap := range ruleDictionary {
			posX, xExists := update.Dictionary[x]
			if !xExists {
				continue
			}
			for y := range ruleMap {
				posY, yExists := update.Dictionary[y]
				if !yExists {
					continue
				}
				if posX >= posY {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}

		if !isValid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	// Reorder the invalid updates with a topological sort
	for _, update := range invalidUpdates {
		validSequence, success := reorderUpdate(update, ruleDictionary)
		if !success {
			fmt.Println("Cycle detected or invalid ordering in update:", update.Sequence)
			continue
		}

		// Find the middle page in the newly sorted update sequence
		middleIndex := len(validSequence) / 2
		middlePage := validSequence[middleIndex]
		sumOfMiddlePages += middlePage
	}

	fmt.Printf("Total sum of middle page numbers: %d\n", sumOfMiddlePages)
}

type Update struct {
	Sequence []int
	Dictionary map[int]int
}

func parseInput(filePath string) (map[int]map[int]bool, []Update) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	ruleDictionary := map[int]map[int]bool{}
	updates := []Update{}

	scanner := bufio.NewScanner(file)
	isRuleSection := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRuleSection = false
			continue
		}

		if isRuleSection {
			rulePair := strings.Split(line, "|")
			if len(rulePair) != 2 {
				fmt.Printf("Invalid rule format: '%s'\n", line)
				continue
			}

			x, errX := strconv.Atoi(rulePair[0])
			y, errY := strconv.Atoi(rulePair[1])
			if errX != nil || errY != nil {
				fmt.Printf("Error converting rule '%s': %v %v\n", line, errX, errY)
				continue
			}

			if _, exists := ruleDictionary[x]; !exists {
				ruleDictionary[x] = make(map[int]bool)
			}

			ruleDictionary[x][y] = true

		} else {
			slice := strings.Split(line, ",")

			updateSlice := []int{}
			updateMap := make(map[int]int)

			for i, s := range slice {
				num, err := strconv.Atoi(s)
				if err != nil {
					fmt.Printf("Error converting page '%s': %v\n", s, err)
					continue
				}

				updateSlice = append(updateSlice, num)
				updateMap[num] = i
			}

			updates = append(updates, Update{
				Sequence: updateSlice,
				Dictionary: updateMap,
			})
		}
	}

	return ruleDictionary, updates
}

func reorderUpdate(update Update, ruleDictionary map[int]map[int]bool) ([]int, bool) {
	// Build a directed acyclic graph
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	nodes := make(map[int]bool)

	// Initialize in-degree and nodes
	for _, page := range update.Sequence {
		inDegree[page] = 0
		nodes[page] = true
	}	

	// Add edges based on the ordering rules
	for x, ruleMap := range ruleDictionary {
		if !nodes[x] {
			continue
		}
		for y := range ruleMap {
			if !nodes[y] {
				continue
			}
			graph[x] = append(graph[x], y)
			inDegree[y]++
		}
	}

	// Topologically sorting the DAG using Kahn's algo
	var queue []int
	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	var sortedSequence []int
	for len(queue) > 0 {
		// Dequeue node
		n := queue[0]
		queue = queue[1:]

		sortedSequence = append(sortedSequence, n)

		// Decrease in-degree of adjacent nodes
		for _, m := range graph[n] {
			inDegree[m]--
			if inDegree[m] == 0 {
				queue = append(queue, m)
			}
		}
	}

	// Check for a cycle
	if len(sortedSequence) != len(nodes) {
		return nil, false
	}

	return sortedSequence, true
}