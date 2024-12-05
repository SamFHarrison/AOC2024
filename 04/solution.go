package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// PLAN
// 1. Create a string for every row, column and diagonal
// 2. Search each string for every occurance of 'XMAS' and 'SAMX'

func parseInput(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed opening file %s", err)
	}
	defer file.Close()

	grid := [][]string{}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	return grid, nil
}

func buildColumnStrings(grid [][]string) []string {
	columns := [][]string{}
	for i := range grid[0] {
		column := []string{}
		for _, v := range grid {
			column = append(column, v[i])
		}
		columns = append(columns, column)
	}

	columnStrings := squashSlices(columns)

	return columnStrings
}

func buildDiagonalStrings(grid [][]string) []string {
	diagonals := [][]string{}

	numOfColumns := len(grid[0])
	numOfRows := len(grid)

	// Order of starting coordinates going SE
	//     1 2 3 4 5
	//     x x x x x
	//   6 x x x x x
	//   7 x x x x x
	//   8 x x x x x
	//   9 x x x x x

	// Then starting coordinates going NE
	//   1 x x x x x
	//   2 x x x x x
	//   3 x x x x x
	//   4 x x x x x
	//   5 x x x x x
	//       6 7 8 9

	// along the top row first going south-east
	for ci := range grid[0] {
		x, y := ci, 0
		diagonal := []string{}
		
		for x < numOfColumns && y < numOfRows {
			diagonal = append(diagonal, grid[y][x])
			x++
			y++
		}
		
		diagonals = append(diagonals, diagonal)
	}

	// then down the rows going south-east
	for ci := range grid[0] {
		if ci == 0 {
			continue
		}

		x, y := ci, 0
		diagonal := []string{}
		
		for x < numOfColumns && y < numOfRows {
			diagonal = append(diagonal, grid[x][y])
			x++
			y++
		}
		
		diagonals = append(diagonals, diagonal)
	}

	// then down the rows agin but going north-east
	for ri := range grid {
		x, y := 0, ri
		diagonal := []string{}
		
		for x < numOfColumns && y >= 0 {
			diagonal = append(diagonal, grid[y][x])
			x++
			y--
		}
		
		diagonals = append(diagonals, diagonal)
	}

	// and finally along the bottom row going north-east
	for ci := range grid[0] {
		if ci == 0 {
			continue
		}

		x, y := ci, numOfRows - 1
		diagonal := []string{}
		
		for x < numOfColumns && y >= 0 {
			diagonal = append(diagonal, grid[y][x])
			x++
			y--
		}
		
		diagonals = append(diagonals, diagonal)
	}

	// First attempt

	// numOfColumns := len(rows)
	// numOfRows := len(rows)

	// for ci := range rows[0] {
	// 	diagonal := []string{}
		// for ri := 0; ri < numOfRows; ri++ {
		// 	if ci + 1 < numOfRows {	
		// 		diagonal = append(diagonal, rows[ri][ci + 1])
		// 		ri++
		// 	} else {
		// 		break
		// 	}
		// }
	// 	diagonals = append(diagonals, diagonal)
	// }

	diagonalStrings := squashSlices(diagonals)

	return diagonalStrings
}

func squashSlices(grid [][]string) []string {
	result := []string{}

	for _, innerSlice := range grid {
		// Join the inner slice into a single string
		squashed := strings.Join(innerSlice, "")
		result = append(result, squashed)
	}

	return result
}

func wordSearch(strings []string) int {
	wordCount := 0
	word := regexp.MustCompile(`XMAS`)
	wordBackwards := regexp.MustCompile(`SAMX`)

	for _, s := range strings {
		instances := word.FindAllString(s, -1)
		instances = append(instances, wordBackwards.FindAllString(s, -1)...)
		wordCount += len(instances)
	}

	return wordCount
}

func main()  {

	fmt.Println("🔄 Parsing input...")
	grid, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
	}
	fmt.Printf("✅ Input parsed\n\n")

	fmt.Println("🔄 Building strings for each row...")
	strings := squashSlices(grid)
	fmt.Printf("✅ Row strings built\n\n")

	fmt.Println("🔄 Building strings for each column...")
	columns := buildColumnStrings(grid)
	strings = append(strings, columns...)
	fmt.Printf("✅ Column strings built\n\n")

	fmt.Println("🔄 Building strings for each diagonal...")
	diagonals := buildDiagonalStrings(grid)
	strings = append(strings, diagonals...)
	fmt.Printf("✅ Diagonal strings built\n\n")
	
	wordCount := wordSearch(strings)
	fmt.Println(wordCount)
}