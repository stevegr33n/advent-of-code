package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const symbols = "!@#$%^&*()/+=-_?,;:[]{}`~£¢∞§¶•ªº™¡–≠"

type numbers map[string]string

func main() {
	file, err := os.Open("./engineSchematic.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//n := numbers{}

	var lines []string
	found := 0
	finalNumber := ""
	var fuckingMatch int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	totalRows := len(lines)
	for row := 0; row < totalRows; row++ {
		for col := 0; col < len(lines[row]); col++ {
			c := string(lines[row][col])
			if _, err := strconv.Atoi(c); err == nil {
				found++
			} else {
				for ; found > 0; found-- {
					finalNumber += string(lines[row][col-found])
				}
				for i := 0; i < len(finalNumber); i++ {
					// check the diagonal values around each number, and determine the final value
					//n[strconv.Itoa(row)+"-"+strconv.Itoa(col-i)] = finalNumber
					if checkDiagonals(row, col-i-1, lines, totalRows-1, len(lines[row])-1) {
						f, _ := strconv.Atoi(finalNumber)
						fmt.Println(fuckingMatch, "+", f, "=", fuckingMatch+f)
						fmt.Println("row", row)
						fuckingMatch += f
					}
				}
				finalNumber = ""
			}
		}
	}
	fmt.Print("=========\n", fuckingMatch, "\n")
}

func checkDiagonals(row int, col int, lines []string, rowLength int, colLength int) bool {
	if row > 0 && col > 0 {
		leftUpper := string(lines[row-1][col-1])
		fmt.Println("leftUpper", leftUpper)
		if strings.Contains(symbols, leftUpper) {
			return true
		}
	}
	if row > 0 && col < colLength {
		rightUpper := string(lines[row-1][col+1])
		fmt.Println("rightUpper", rightUpper)
		if strings.Contains(symbols, rightUpper) {
			return true
		}
	}
	if row < rowLength && col > 0 {
		leftLower := string(lines[row+1][col-1])
		fmt.Println("leftLower", leftLower)
		if strings.Contains(symbols, leftLower) {
			return true
		}
	}
	if row < rowLength && col < colLength {
		rightLower := string(lines[row+1][col+1])
		fmt.Println("rightLower", rightLower)
		if strings.Contains(symbols, rightLower) {
			return true
		}
	}
	// this only needs to happen on the first and last digits, but whatever
	if col > 0 {
		left := string(lines[row][col-1])
		if strings.Contains(symbols, left) {
			return true
		}
	}
	if col < colLength {
		right := string(lines[row][col+1])
		if strings.Contains(symbols, right) {
			return true
		}
	}
	return false
}
