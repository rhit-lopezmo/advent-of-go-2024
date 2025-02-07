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
func parseLine(line string, vals *[]int) error {
	// trim line to remove new line
	actualLine :=	strings.TrimSuffix(line, "\n")

	// split the line by spaces
	valStrs := strings.Fields(actualLine)	
	
	// convert each value in the line and append to array
	for _, valStr := range valStrs {
		// convert
		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Printf("error converting string to int: %s\n%d\n", err, val)
			return err
		}

		// append
		*vals = append(*vals, val)
	}	

	// no error
	return nil
}

func main() {
	// open file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a reader
	reader := bufio.NewReader(file)

	// create vals
	vals := []int{}

	// count how many lines were read
	numLines := 0

	// read line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				numLines++
				err = parseLine(line, &vals) // parse line into values
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
		err = parseLine(line, &vals)
		if err != nil {
			fmt.Println("error converting string to int:", err)
			return
		}
		numLines++
	}

	// sort the vals in increasing order
	sort.Ints(vals)

	fmt.Printf("is multiple of 2?: %t", numLines * 2 % 2 == 0)
	
}
