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
	var letter string
	var combinedNumber string
	var j int

	for j = 0; j < len(line); j++ {
		letter = string(line[j])
		if _, err := strconv.Atoi(letter); err == nil {
			combinedNumber += letter
			break
		}
	}
	for k := len(line) - 1; k > j; k-- {
		letter = string(line[k])
		if _, err := strconv.Atoi(letter); err == nil {
			combinedNumber += letter
			break
		}
	}

	if len(combinedNumber) == 1 {
		combinedNumber += combinedNumber
	}

	val, err := strconv.Atoi(combinedNumber)
	if err != nil {
		os.Exit(1)
	}
	return val
}
