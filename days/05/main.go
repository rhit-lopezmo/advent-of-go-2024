package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StringSet map[string]struct{}

func (set StringSet) add(value string) {
	set[value] = struct{}{}
}

func (set StringSet) remove(value string) {
	delete(set, value)
}

func (set StringSet) contains(value string) bool {
	_, exists := set[value]
	
	return exists
}

func (set StringSet) toString() string {
	str := "["

	for key := range set {
		str = str + key + ", "
	}

	if len(str) > 1 {
		str = str[:len(str) - 2]
	}

	str = str + "]"

	return str 
}


func processRule(rule string, ruleMap map[string][]string) {
	// split the rule
	data := strings.Split(rule, "|")

	// separate into what needs to come before + after
	before := data[0]
	after := data[1]


	// append if it already exists, otherwise make a new entry with before as the key + after as first val
	_, exists := ruleMap[before]
	if (exists) {
		ruleMap[before] = append(ruleMap[before], after)
	} else {
		ruleMap[before] = []string{after}
	}
}

func isValidUpdate(data []string, ruleMap map[string][]string) bool {
	// track what's already been checked so we know what to check the ruleMap for
	checked := StringSet{}

	// check if each entry is valid
	for _, entry := range data {
		_, existInRules := ruleMap[entry]

		if existInRules {
			for _, numAfter := range ruleMap[entry] {
				if checked.contains(numAfter) {
					return false
				}
			}
		}

		checked.add(entry)
	}


	// all entries were valid
	return true
}

func processUpdate(update string, ruleMap map[string][]string, validUpdateMidNums *[]string) {
	data := strings.Split(update, ",") // split the update
	valid := isValidUpdate(data, ruleMap)

	// find middle number
	midNum := data[len(data) / 2]

	if valid {
		*validUpdateMidNums = append(*validUpdateMidNums, midNum)
	}
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

	// track if in rules or updates
	inRules := true

	// map for all the rules in the input
	ruleMap := make(map[string][]string)
	
	// tracks all valid mid nums
	validMidNums := []string{}

	// read line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Printf("Error reading from file: %s\n", err)
				return
			}
		}

		// get rid of newline
		line = strings.TrimSuffix(line, "\n")

		// swapping from rules to updates, just continue
		if (len(line) == 0) {
			inRules = false
			continue
		}
	
		// process as a rule or as an update
		if inRules {
			processRule(line, ruleMap)
		} else {
			processUpdate(line, ruleMap, &validMidNums)
		}
	}
	
	// calculate the sum
	sum := 0
	for _, numStr := range validMidNums {
		num, err := strconv.Atoi(numStr)		
		if err != nil {
			fmt.Printf("Could not convert '%s' to an int.\n", numStr)
			return
		}

		sum += num
	}

	fmt.Printf("total sum of valid mid nums: %d\n", sum)
}
