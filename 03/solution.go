package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed opening file %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var resultBuilder strings.Builder
	for scanner.Scan() {
		resultBuilder.WriteString(scanner.Text())
	}

	return resultBuilder.String(), nil
}

// func multiplyPairs(strings []string) (int, error) {
// 	sum := 0

// 	regex := regexp.MustCompile(`-?\d+`)

// 	for _, string := range strings {
// 		pair := regex.FindAllString(string, -1)

// 		num1, err := strconv.Atoi(pair[0])
// 		if err != nil {
// 			fmt.Printf("Error converting '%s' to integer: %v\n",pair[0], err)
// 		}
// 		num2, err := strconv.Atoi(pair[1])
// 		if err != nil {
// 			fmt.Printf("Error converting '%s' to integer: %v\n",pair[1], err)
// 		}

// 		sum += num1 * num2
// 	}

// 	return sum, nil
// }

// F in the chat for my first attempt
// it doesn't take into account that the do's and dont's
// aren't guarunteed to be perfectly alternating

// func removeDisabledMuls(input string) string {
//     patternDont := `don't\(\)`
//     patternDo := `do\(\)`

//     dontRegex := regexp.MustCompile(patternDont)
//     doRegex := regexp.MustCompile(patternDo)

//     var resultBuilder strings.Builder
//     currentIndex := 0

//     dontIndices := dontRegex.FindAllStringIndex(input, -1)
// 	fmt.Println(dontIndices)

//     for _, dontPos := range dontIndices {
//         dontStart := dontPos[0]
//         dontEnd := dontPos[1]

// 		fmt.Printf("currentIndex: %d, dontStart: %d\n", currentIndex, dontStart)
//         resultBuilder.WriteString(input[currentIndex:dontStart])

//         nextDoPos := doRegex.FindStringIndex(input[dontEnd:])
		
//         if nextDoPos != nil {
// 			doEnd := dontEnd + nextDoPos[1]
//             currentIndex = doEnd
// 			fmt.Println(doEnd)
// 		} else {
// 			currentIndex = len(input)
// 			break
// 		}
// 	}

//     resultBuilder.WriteString(input[currentIndex:])

//     return resultBuilder.String()
// }

type Event struct {
	Kind   string   // "mul", "do", or "don't"
	Start  int      // Start index of the event
	End    int      // End index of the event
	Groups []string // Captured groups (for mul)
}

func createEvents(input string) []Event {
	patternMul := `mul\((-?\d+),(-?\d+)\)`
	patternDo := `do\(\)`
	patternDont := `don't\(\)`
	
	var events []Event
	
	// Find all mul(x,y) instances
	mulRegex := regexp.MustCompile(patternMul)
	mulMatches := mulRegex.FindAllStringSubmatchIndex(input, -1)
	for _, m := range mulMatches {
		events = append(events, Event{
			Kind:   "mul",
			Start:  m[0],
			End:    m[1],
			Groups: []string{input[m[2]:m[3]], input[m[4]:m[5]]},
		})
	}
	
	// Find all do() instances
	doRegex := regexp.MustCompile(patternDo)
	doMatches := doRegex.FindAllStringIndex(input, -1)
	for _, m := range doMatches {
		events = append(events, Event{
			Kind:  "do",
			Start: m[0],
			End:   m[1],
		})
	}
	
	// Find all don't() instances
	dontRegex := regexp.MustCompile(patternDont)
	dontMatches := dontRegex.FindAllStringIndex(input, -1)
	for _, m := range dontMatches {
		events = append(events, Event{
			Kind:  "don't",
			Start: m[0],
			End:   m[1],
		})
	}

	return events
}

func multiplyPairs(events []Event) int {
	enabled := true
	total := 0

	for _, event := range events {
		switch event.Kind {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				xStr := event.Groups[0]
				yStr := event.Groups[1]
				x, err := strconv.Atoi(xStr)
				if err != nil {
					fmt.Printf("Error parsing x: %v\n", err)
					continue
				}
				y, err := strconv.Atoi(yStr)
				if err != nil {
					fmt.Printf("Error parsing y: %v\n", err)
					continue
				}
				result := x * y
				total += result
				fmt.Printf("mul(%d,%d) = %d\n", x, y, result)
			} else {
				fmt.Printf("Skipped mul(%s,%s) due to disable\n", event.Groups[0], event.Groups[1])
			}
		}
	}

	return total
}

func main() {
	fmt.Println("🔄 Parsing input...")
	input, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
	}
	fmt.Printf("☑️  Input parsed\n\n")
	
	// Part 2
	fmt.Println("🔄 Create event timeline of string input...")
	events := createEvents(input)
	fmt.Printf("☑️  Event timeline created successfully\n\n")

	// Part of my first attempt
	// fmt.Println("🔄 Finding all instances of 'mul(x,y)'...")
	// regex := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)`)
	// regexMatches := regex.FindAllString(input, -1)
	// fmt.Printf("☑️  Found all instances\n\n")

	fmt.Println("🔄 Sorting events into the order they occur...")
	sort.Slice(events, func(i, j int) bool {
		return events[i].Start < events[j].Start
	})
	fmt.Printf("☑️  Events sorted successfully\n\n")
	

	fmt.Println("🔄 Multiplying pairs...")
	sum := multiplyPairs(events)
	if err != nil {
		fmt.Printf("Error multiplying pairs: %s\n", err)
	}
	fmt.Printf("☑️  Pairs multiplied\n\n")
	
	fmt.Println("Total sum of all pairs:", sum)	
}