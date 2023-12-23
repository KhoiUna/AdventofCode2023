package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ANS := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		winningNums := formatInput(strings.Split(strings.Split(strings.Split(line, "|")[0], ":")[1], " "))
		haveNums := formatInput(strings.Split(strings.Split(line, "|")[1], " "))

		fmt.Printf("%#v\n", winningNums)
		fmt.Printf("%#v\n", haveNums)

		count := 0
		for _, num := range haveNums {
			if !slices.Contains(winningNums, num) {
				continue
			}

			if count == 0 {
				count++
			} else {
				count *= 2
			}
		}
		fmt.Print("count->", count, "\n")
		ANS += count
	}

	fmt.Println("ANS:", ANS)
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
