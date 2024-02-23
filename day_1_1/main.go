package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./numbers.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res int
	for scanner.Scan() {
		val := calibrationValues(scanner.Text())
		fmt.Print(res, " + ", val, " = ")
		res += val
		fmt.Print(res, "\n")
	}
	fmt.Print("=========\n", res, "\n")
}

func calibrationValues(line string) int {
	var j int
	var result string
	var char string

	for j = 0; j < len(line); j++ {
		char = string(line[j])
		if _, err := strconv.Atoi(char); err == nil {
			result += char
			break
		}
	}
	for k := len(line) - 1; k > j; k-- {
		char = string(line[k])
		if _, err := strconv.Atoi(char); err == nil {
			result += char
			break
		}
	}

	if len(result) == 1 {
		result += result
	}

	integerResult, err := strconv.Atoi(result)
	if err != nil {
		os.Exit(1)
	}
	return integerResult
}
