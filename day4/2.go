package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	rows := 0
	for scanner.Scan() {
		rows++
		lines = append(lines, scanner.Text())
	}
	scratches := []int{}
	for i := 0; i < rows; i++ {
		scratches = append(scratches, 1)
	}

	for _, line := range lines {
		currentCard := strings.TrimSpace(strings.Split(strings.Split(line, "Card")[1], ":")[0])
		currentCardNum, _ := strconv.Atoi(currentCard)
		winningNums := formatInput(strings.Split(strings.Split(strings.Split(line, "|")[0], ":")[1], " "))
		haveNums := formatInput(strings.Split(strings.Split(line, "|")[1], " "))

		copies := scratches[currentCardNum-1]
		for i := 0; i < findNumOfMatches(winningNums, haveNums); i++ {
			scratches[currentCardNum+i] += copies
		}
	}

	ANS := 0
	for _, num := range scratches {
		ANS += num
	}
	fmt.Println("->ANS:", ANS)
}

func formatInput(input []string) []string {
	newInput := []string{}
	for _, val := range input {
		if strings.TrimSpace(val) != "" {
			newInput = append(newInput, strings.TrimSpace(val))
		}
	}
	return newInput
}

func findNumOfMatches(winningNums, haveNums []string) int {
	count := 0
	for _, num := range winningNums {
		if !slices.Contains(haveNums, num) {
			continue
		}
		count++
	}
	return count
}
