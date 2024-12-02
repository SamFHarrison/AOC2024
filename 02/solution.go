package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([][]int, error) {
	reports := [][]int{}
	
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed opening file %s", err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		report := []int{}

		for _, value := range parts {
			int, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("failed converting '%s' to an integer %v", value, err)
			}

			report = append(report, int)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func isReportSafe(report []int) bool {
    if len(report) < 2 {
        return true
    }

    firstDifference := report[1] - report[0]
    if firstDifference == 0 {
        return false
    }

    direction := ""
    if firstDifference > 0 {
        direction = "increasing"
    } else {
        direction = "decreasing"
    }

    for i := 1; i < len(report); i++ {
        difference := report[i] - report[i - 1]

        if difference == 0 {
            return false
        }

        if absoluteValue(difference) < 1 || absoluteValue(difference) > 3 {
            return false
        }

        if direction == "increasing" && difference < 0 {
            return false
        }
        if direction == "decreasing" && difference > 0 {
            return false
        }
    }

    return true
}

func absoluteValue(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func isReportSafeWithDampener(report []int) bool {
	for i := 0; i < len(report); i++ {

		// I'm removing the current report[i] by creating a new slice
		// with all the values before and after the current index
		// excluding the current index itself. I'm learning so many
		// little tricks already and it's only day 2 of using Go
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i + 1:]...)

		if isReportSafe(modifiedReport) {
			return true
		}
	}

	return false
}

func main() {

	fmt.Println("🔄 Parsing input...")
    reports, err := parseInput("input.txt")
    if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
    }
	fmt.Printf("☑️  Input parsed\n\n")

	fmt.Println("🔄 Starting program...")
    safeReports := 0
    for _, report := range reports {
        if isReportSafe(report) {
            safeReports++
		} else if isReportSafeWithDampener(report) {
            safeReports++
		}
    }
	fmt.Printf("☑️  Program complete\n\n")

    fmt.Println("Total safe reports:", safeReports)
}
