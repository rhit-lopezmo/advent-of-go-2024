package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const NUM_MATCHES = -1 // finds all possible matches in regex

func main() {
	// get args from program call
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("usage: go run main.go <filename>\n")
		return
	}

	// get filename
	filename := args[1]

	// read entire file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}
	
	// convert data to a string
	content := string(data)

	// (regex) template to find mul expressions
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`) // capture the numbers from the mul expression

	// get matches
	matches := re.FindAllStringSubmatch(content, NUM_MATCHES)

	// calculate total
	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		sum += num1 * num2
	}
		
	fmt.Printf("total from mul sum: %d\n", sum)
}
