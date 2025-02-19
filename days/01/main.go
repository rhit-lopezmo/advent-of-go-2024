package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// parses a line and stores the two values into the array
func parseLine(line string, leftVals *[]int, rightVals *[]int) error {
	// trim line to remove new line
	actualLine :=	strings.TrimSuffix(line, "\n")

	// split the line by spaces
	valStrs := strings.Fields(actualLine)	
	
	// convert each value in the line and append to array
	for i, valStr := range valStrs {
		// convert
		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Printf("error converting string to int: %s\n%d\n", err, val)
			return err
		}

		// append to left list if 0, otherwise right list
		if i == 0 {
			*leftVals = append(*leftVals, val)
		} else {
			*rightVals = append(*rightVals, val)
		}
	}	

	// no error
	return nil
}

// absolute value function for integers
func absInt(value int) int {
	if value < 0 {
		return -1 * value
	} else {
		return value
	}
}

func main() {
	// get args
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("usage: go run main.go <filename>\n")
		return
	}

	// get filename
	filename := args[1] // args[0] is program name like C

	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a reader
	reader := bufio.NewReader(file)

	// create vals
	leftVals := []int{}
	rightVals := []int{}

	// count how many lines were read
	numLines := 0

	// read line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				numLines++
				err = parseLine(line, &leftVals, &rightVals) // parse line into values
				if err != nil {
					return
				}
				break;
			} else {
				fmt.Println("Read err:", err)
				return;
			}
		}
	
		// parse the line
		err = parseLine(line, &leftVals, &rightVals)
		if err != nil {
			fmt.Println("error converting string to int:", err)
			return
		}
		numLines++
	}

	// sort the vals in increasing order
	sort.Ints(leftVals)
	sort.Ints(rightVals)

	// find total sum of distances
	sum := 0

	for i := range leftVals {
		sum += absInt(leftVals[i] - rightVals[i])
	}

	fmt.Printf("total distance sum: %d\n", sum)
}
