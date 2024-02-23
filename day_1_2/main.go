package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numberLookup(chars string) (string, error) {
	numbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for k, v := range numbers {
		if strings.Contains(chars, k) {
			return v, nil
		}
	}
	return "", errors.New("no match")
}

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
	var chars string

	for j = 0; j < len(line); j++ {
		char = string(line[j])
		if _, err := strconv.Atoi(char); err == nil {
			result = char
			chars = ""
			break
		}
		chars += char
		if match, err := numberLookup(chars); err == nil {
			result = match
			chars = ""
			break
		}
	}
	for k := len(line) - 1; k > j; k-- {
		char = string(line[k])
		if _, err := strconv.Atoi(char); err == nil {
			result += char
			break
		}
		chars = char + chars
		if match, err := numberLookup(chars); err == nil {
			result += match
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
