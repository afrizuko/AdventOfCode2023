package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var config = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {

	f, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var totalGames int
	var fewest int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.SplitN(line, ": ", 2)
		gameSets := strings.Split(game[1], ";")

		result, err := partTwo(gameSets)
		if err == nil {
			fewest += result
		}

		if partOne(gameSets) == nil {
			gameID, err := strconv.Atoi(strings.Split(game[0], " ")[1])
			if err == nil {
				totalGames += gameID
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println("Part 1:", totalGames)
	fmt.Println("Part 2:", fewest)
}

func partOne(gameSets []string) error {
	for _, gameSet := range gameSets {
		for _, part := range strings.Split(gameSet, ",") {
			entries := strings.Split(strings.Trim(part, " "), " ")
			color := strings.Trim(entries[1], " ")

			entry, err := strconv.Atoi(entries[0])
			if err == nil && entry > config[color] {
				return io.EOF
			}
		}
	}
	return nil
}

func partTwo(gameSets []string) (int, error) {
	var minRequired = make(map[string]int)
	for _, gameSet := range gameSets {
		for _, part := range strings.Split(gameSet, ",") {
			entries := strings.Split(strings.Trim(part, " "), " ")
			color := strings.Trim(entries[1], " ")

			entry, err := strconv.Atoi(entries[0])
			if err != nil {
				return 0, err
			}
			minRequired[color] = max(entry, minRequired[color])
		}
	}

	return minRequired["red"] * minRequired["blue"] * minRequired["green"], nil
}
