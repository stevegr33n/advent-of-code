package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type allCharacters map[string]string
type numberCoordinates map[string]string

func main() {
	file, err := os.Open("./engineSchematic.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	all := allCharacters{}
	n := numberCoordinates{}

	//var res int
	var lineNum int

	for scanner.Scan() {
		_ = processSchematic(scanner.Text(), all, n, lineNum)
		lineNum++
		//fmt.Print(res, " + ", val, " = ")
		//res += val
		//fmt.Print(res, "\n")
	}
	fmt.Print("=========\n", all, "\n")
	fmt.Print("=========\n", n, "\n")
	//fmt.Print("=========\n", res, "\n")
}

func processSchematic(line string, all allCharacters, n numberCoordinates, lineNum int) int {
	for rowNum, char := range line {
		c := string(char)
		all[strconv.Itoa(lineNum)+strconv.Itoa(rowNum)] = c
		if _, err := strconv.Atoi(c); err == nil {
			// if its a number then store the coordinates
			n[strconv.Itoa(lineNum)+strconv.Itoa(rowNum)] = c
		}
		if all[lineNum] {
		}
	}
	return 0
}
