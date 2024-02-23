package main

import (
	"bufio"
	"fmt"
	"os"
)

func getCubes(color string, number int) bool {
	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

}

func main() {
	file, err := os.Open("./games.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res int
	for scanner.Scan() {
		val := gameID(scanner.Text())
		fmt.Print(res, " + ", val, " = ")
		res += val
		fmt.Print(res, "\n")
	}
	fmt.Print("=========\n", res, "\n")
}

func gameID(line string) int {

	return 0
}
