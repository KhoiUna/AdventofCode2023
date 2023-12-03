package main

import (
	"bufio"
	"strconv"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
leftDigit, rightDigit := findDigits(line)
		sum += formatResult(leftDigit, rightDigit)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
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
		return right * 10 + right
	} else {
		return left*10 + right
	}
}
