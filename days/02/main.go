package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// gets the difference between two values
func diff(val1 int, val2 int) int {
	result := val1 - val2

	if result < 0 {
		return -1 * result
	} else {
		return result
	}
}

// use this strategy if first number < second number
func checkDec(nums []int) bool {
	// check if strictly decreasing
	for i := 0; i + 1 < len(nums); i++ {
		if nums[i] >= nums[i + 1] || diff(nums[i], nums[i + 1]) > 3 {
			return false
		}
	}

	return true
}

// use this strategy if first number < second number
func checkInc(nums []int) bool {
	// check if strictly increasing
	for i := 0; i + 1 < len(nums); i++ {
		if nums[i] <= nums[i + 1] || diff(nums[i], nums[i + 1]) > 3 {
			return false
		}
	}

	return true
}

func reportSafety(nums []int) bool {
	if len(nums) == 0 {
		fmt.Printf("Empty line found, returning true from reportSafety\n")
		return true
	}

	if nums[0] < nums[1] {
		return checkDec(nums)
	} else if nums[0] > nums[1] {
		return checkInc(nums)
	} else {
		return false
	}
}

func checkReport(reportStr string) (bool, error) {
	// parse out data from line
	numsStrs := strings.Fields(reportStr)
	nums := []int{}
	for _, numStr := range numsStrs {
		// convert
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("error converting string to int: %s\n%s\n", err, numStr)
			return false, err
		}
		
		// append it to the array
		nums = append(nums, num)
	}

	return reportSafety(nums), nil
}

func main() {
	// get args from program call
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("usage: go run main.go <filename>\n")
		return
	}

	// get filename
	filename := args[1]

	// open file for reading
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening '%s': %s\n", filename, err)
	}

	// create a reader for the file
	reader := bufio.NewReader(file)

	// track number of safe reports read
	safeReports := 0

	// read the line by line and figure out if unsafe
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				// check if empty line, if so just break
				if len(line) == 0 {
					break
				}

				// check safety
				actualLine := strings.TrimSuffix(line, "\n")	
				isSafe, err := checkReport(actualLine)
				if err != nil {
					return
				}

				if isSafe {
					safeReports++
				}
				break
			} else {
				fmt.Printf("Error reading line from file: %s\n", err)
				return
			}
		}

		// determine if unsafe report and add it
		actualLine := strings.TrimSuffix(line, "\n")
		isSafe, err := checkReport(actualLine)
		if err != nil {
			return
		}

		if isSafe {
			safeReports++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeReports)
}
