package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./games.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res int
	for scanner.Scan() {
		val := processGames(scanner.Text())
		fmt.Print(res, " + ", val, " = ")
		res += val
		fmt.Print(res, "\n")
	}
	fmt.Print("=========\n", res, "\n")
}

// example input
// Game 1: 5 red, 1 green; 6 red, 3 blue; 9 red; 1 blue, 1 green, 4 red; 1 green, 2 blue; 2 blue, 1 red
func gameIsPossible(game string) bool {
	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	rolls := strings.Split(game, "; ")

	for _, roll := range rolls {
		dice := strings.Split(roll, ", ")
		for _, die := range dice {
			split := strings.Split(die, " ")
			number, colour := split[0], split[1]
			num, err := strconv.Atoi(number)
			if err != nil {
				os.Exit(1)
			}
			if num > cubes[colour] {
				return false
			}
		}
	}
	return true
}

func extractGameID(line string) (int, string) {
	split := strings.Split(line, ": ")
	ID, game := split[0][5:], split[1]
	numericalID, err := strconv.Atoi(ID)
	if err != nil {
		return 0, ""
	}
	return numericalID, game
}

func processGames(line string) int {
	ID, game := extractGameID(line)
	if gameIsPossible(game) {
		return ID
	}
	return 0
}
