package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([]int, []int, error) {

	// Initilaise my slices
	list1, list2 := []int{}, []int{} 

	// Read input from txt file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed opening file %s", err)
	}
	defer file.Close()

	// Reading file 1 line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Splitting the string at the whitespace in between the numbers
		parts := strings.Fields(scanner.Text())

		// converting both numbers from strings to ints
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed converting '%s' to an integer %v", parts[0], err)
		}
		list1 = append(list1, num1)

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting '%s' to an integer: %v", parts[1], err)
		}
		list2 = append(list2, num2)
	}

	return list1, list2, nil
}

func mergeSort(slice []int) []int {

	// Stop recursion when slice has <= 1 ints left
	if len(slice) <= 1 {
		return slice
	}

	// Get the middle point of the slice
	middle := len(slice) / 2

	// Split my slice in half
	left := mergeSort(slice[:middle])
	right := mergeSort(slice[middle:])

	return merge(left, right)
}

func merge(leftSlice, rightSlice []int) []int {
	
	result := []int{}
	leftPointer, rightPointer := 0, 0

	// Loop over the slices
	for leftPointer < len(leftSlice) && rightPointer < len(rightSlice) {
		
		// Add the smaller value to result slice and move that pointer forwards by 1
		if leftSlice[leftPointer] < rightSlice[rightPointer] {
			result = append(result, leftSlice[leftPointer])
			leftPointer++
		} else {
			result = append(result, rightSlice[rightPointer])
			rightPointer++
		}
	}

	// Adding any excess values to result slice in case the lists are different lengths
	for leftPointer < len(leftSlice) {
		result = append(result, leftSlice[leftPointer])
		leftPointer++
	}
	for rightPointer < len(rightSlice) {
		result = append(result, rightSlice[rightPointer])
		rightPointer++
	}

	return result
}

func findSimilarityScore(leftSlice, rightSlice []int) int {
	leftPointer, rightPointer := 0, 0
	similarityScore := 0

	for leftPointer < len(leftSlice) && rightPointer < len(rightSlice) {
		if leftSlice[leftPointer] == rightSlice[rightPointer] {
			count := 1

			for rightPointer + 1 <= len(rightSlice) && rightSlice[rightPointer + 1] == rightSlice[rightPointer] {
				count++
				rightPointer++
			}

			similarityScore += leftSlice[leftPointer] * count
			leftPointer++

		} else if leftSlice[leftPointer] < rightSlice[rightPointer] {
			leftPointer++
		} else {
			rightPointer++
		}
	}

	return similarityScore
}

func main() {

	fmt.Println("Starting part 1...")

	list1, list2, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
	}

	sortedList1 := mergeSort(list1)
	sortedList2 := mergeSort(list2)

	differences := []int{}
	list1Pointer, list2Pointer := 0, 0

	for list1Pointer < len(sortedList1) && list2Pointer < len(sortedList2) {
		difference := sortedList1[list1Pointer] - sortedList2[list2Pointer]
		if difference < 0 {
			difference = -difference
		}

		differences = append(differences, difference)

		list1Pointer++
		list2Pointer++
	}

	sumOfDifferences := 0

	for _, value := range differences {
		sumOfDifferences += value
	}

	fmt.Println("Sum of Differences:", sumOfDifferences)

	fmt.Println("Starting part 2...")

	similarityScore := findSimilarityScore(sortedList1, sortedList2)

	fmt.Println("Similiarity Score:", similarityScore)	
}