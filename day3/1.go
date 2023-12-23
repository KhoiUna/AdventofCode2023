package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
1. Loop each char in each line
2. Get the start and end position of the digit
3. Loop through neighbors of that digit and see if there's a symbol
   -> if there's a symbol -> append to parts
4. Sum of parts
*/

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	matrix := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	digit := ""
	parts := []int{}

	for rowId := range matrix {
		start := -1

		for colId := range matrix[rowId] {
			char := matrix[rowId][colId]

			if isDigit(rune(char)) {
				if digit == "" {
					start = colId
				}
				digit += string(char)

				// Last char in row
				if colId == len(matrix[0])-1 {
					if digit != "" {
						if isPart(start, start+len(digit)-1, rowId, matrix) {
							num, _ := strconv.Atoi(digit)
							parts = append(parts, num)
						}
					}
					// Reset `digit`
					digit = ""
				}
			} else {
				if digit != "" {
					if isPart(start, start+len(digit)-1, rowId, matrix) {
						num, _ := strconv.Atoi(digit)
						parts = append(parts, num)
					}
				}
				// Reset `digit`
				digit = ""
			}
		}
	}

	ans := 0
	for _, num := range parts {
		ans += num
	}
	fmt.Println("ANS:", ans)
}

func isDigit(input rune) bool {
	return '0' <= input && input <= '9'
}

func isSymbol(input rune) bool {
	return !isDigit(input) && input != '.'
}

func isPart(startIdx, endIdx, rowId int, matrix []string) bool {
	if startIdx > 0 {
		startIdx--
		if isSymbol(rune(matrix[rowId][startIdx])) {
			return true
		}
	}
	if endIdx < len(matrix[rowId])-1 {
		endIdx++
		if isSymbol(rune(matrix[rowId][endIdx])) {
			return true
		}
	}
	for start := startIdx; start <= endIdx; start++ {
		// Up
		if rowId > 0 && isSymbol(rune(matrix[rowId-1][start])) {
			return true
		}
		// Down
		if rowId < len(matrix)-1 && isSymbol(rune(matrix[rowId+1][start])) {
			return true
		}
	}
	return false
}
