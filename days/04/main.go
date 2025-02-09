package main

import (
	"bufio"
	"fmt"
	"os"
)

// Helper function to check if a position is within bounds
func isValid(grid [][]byte, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}

// Function to check a string in a given direction
func checkVerticalUp(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row - 3, col) {
		return 0
	}

	// Collect characters from (row, col) upwards
	word := string([]byte{
		grid[row][col],
		grid[row - 1][col],
		grid[row - 2][col],
		grid[row - 3][col],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

func checkVerticalDown(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row + 3, col) {
		return 0
	}

	word := string([]byte{
		grid[row][col],
		grid[row + 1][col],
		grid[row + 2][col],
		grid[row + 3][col],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

func checkHorizontalLeft(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row, col - 3) {
		return 0
	}

	word := string([]byte{
		grid[row][col],
		grid[row][col - 1],
		grid[row][col - 2],
		grid[row][col - 3],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

func checkHorizontalRight(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row, col + 3) {
		return 0
	}

	word := string([]byte{
		grid[row][col],
		grid[row][col + 1],
		grid[row][col + 2],
		grid[row][col + 3],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

func checkDiagonalUpLeft(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row - 3, col - 3) {
		return 0
	}

	word := string([]byte{
		grid[row][col],
		grid[row - 1][col - 1],
		grid[row - 2][col - 2],
		grid[row - 3][col - 3],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

func checkDiagonalUpRight(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row - 3, col + 3) {
		return 0
	}

	word := string([]byte{
		grid[row][col],
		grid[row - 1][col + 1],
		grid[row - 2][col + 2],
		grid[row - 3][col + 3],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

func checkDiagonalDownLeft(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row + 3, col - 3) {
		return 0
	}

	word := string([]byte{
		grid[row][col],
		grid[row + 1][col - 1],
		grid[row + 2][col - 2],
		grid[row + 3][col - 3],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

func checkDiagonalDownRight(grid [][]byte, row, col int, target string) int {
	if !isValid(grid, row, col) || !isValid(grid, row + 3, col + 3) {
		return 0
	}

	word := string([]byte{
		grid[row][col],
		grid[row + 1][col + 1],
		grid[row + 2][col + 2],
		grid[row + 3][col + 3],
	})

	isXMAS := word == target

	if (isXMAS) {
		return 1
	} else {
		return 0
	}
}

// finds all possible XMAS's (vertical, horizontal, and diagonal)
func findAllXMAS(wordSearch [][]byte) int {
	total := 0

	for i, line := range wordSearch	{
		for j, char := range line {
			if char != 'X' {
				continue;
			}

			total += checkDiagonalDownLeft(wordSearch, i, j, "XMAS")
			total += checkDiagonalDownRight(wordSearch, i, j, "XMAS")
			total += checkDiagonalUpLeft(wordSearch, i, j, "XMAS")
			total += checkDiagonalUpRight(wordSearch, i, j, "XMAS")
			total += checkHorizontalLeft(wordSearch, i, j, "XMAS")
			total += checkHorizontalRight(wordSearch, i, j, "XMAS")
			total += checkVerticalDown(wordSearch, i, j, "XMAS")
			total += checkVerticalUp(wordSearch, i, j, "XMAS")
		}
	}

	return total
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

	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// setup reader
	reader := bufio.NewReader(file)

	// read line by line and add into byte 2d slice
	wordSearch := [][]byte{}
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Printf("Error reading from file: %s\n", err)
				return
			}
		}

		// append everything except newline
		wordSearch = append(wordSearch, line[:len(line) - 1])
	}

	// process wordSearch to find all possible XMAS's
	fmt.Printf("num of XMAS's: %d\n", findAllXMAS(wordSearch))
}
