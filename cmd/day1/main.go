package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var mapper = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var total1, total2 int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text
		result1 := partOne(line())

		if len(result1) > 0 {
			num, err := strconv.Atoi(string(result1[0]) + string(result1[len(result1)-1]))
			if err == nil {
				total1 += num
			}
		}

		result2 := partTwo(line())
		if len(result2) > 0 {
			num, err := strconv.Atoi(string(result2[0]) + string(result2[len(result2)-1]))
			if err == nil {
				total2 += num
			}
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d", total1, total2)
}

func partOne(line string) string {
	var pairs strings.Builder
	for _, l := range line {
		if unicode.IsDigit(l) {
			pairs.WriteRune(l)
		}
	}
	return pairs.String()
}

func partTwo(line string) string {
	var builder, pairs strings.Builder
	for _, l := range line {

		if unicode.IsLetter(l) {
			builder.WriteRune(l)
			for k, v := range mapper {
				if strings.Contains(builder.String(), k) {
					builder.Reset()
					builder.WriteRune(l)
					pairs.WriteString(strconv.Itoa(v))
				}
			}
		}
		if unicode.IsDigit(l) {
			pairs.WriteRune(l)
			builder.Reset()
		}
	}

	return pairs.String()
}
