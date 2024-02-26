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

type cubes map[string]int

// example input
// Game 1: 5 red, 1 green; 6 red, 3 blue; 9 red; 1 blue, 1 green, 4 red; 1 green, 2 blue; 2 blue, 1 red
func calculateMinimumNumberOfCubesForEachGame(game string) cubes {
	c := cubes{
		"red":   0,
		"green": 0,
		"blue":  0,
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
			if num > c[colour] {
				c[colour] = num
			}
		}
	}
	return c
}

func removeGameID(line string) string {
	split := strings.Split(line, ": ")
	return split[1]
}

func processGames(line string) int {
	game := removeGameID(line)
	c := calculateMinimumNumberOfCubesForEachGame(game)

	return c["red"] * c["blue"] * c["green"]
}
