package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var wordToNum = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sumPartOne := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		leftDigit, rightDigit := findDigits(line)
		sumPartOne += formatResult(leftDigit, rightDigit)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumPartOne)
}

func findDigits(line string) (rune, rune) {
	var leftDigit, rightDigit rune
	foundLeft := false
	for _, r := range line {
		if unicode.IsDigit(r) {
			if !foundLeft {
				leftDigit = r
				foundLeft = true
			}
			rightDigit = r
		}
	}
	return leftDigit, rightDigit
}

func formatResult(leftDigit, rightDigit rune) int {
	left, _ := strconv.Atoi(string(leftDigit))
	right, _ := strconv.Atoi(string(rightDigit))

	if left == 0 {
		return right*10 + right
	} else {
		return left*10 + right
	}
}
