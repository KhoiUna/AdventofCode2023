package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	parts = []int{}
	gears = [][][]int{}
)

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

	rows := len(matrix)
	cols := len(matrix[0])
	gears = make([][][]int, rows)
	for i := range gears {
		gears[i] = make([][]int, cols)
	}
	digit := ""

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
						isPart(start, start+len(digit)-1, digit, rowId, matrix)
					}
					// Reset `digit`
					digit = ""
				}
			} else {
				if digit != "" {
					isPart(start, start+len(digit)-1, digit, rowId, matrix)
				}
				// Reset `digit`
				digit = ""
			}
		}
	}

	ans := 0
	for rowId := range matrix {
		for colId := range matrix[rowId] {
			nums := gears[rowId][colId]
			if matrix[rowId][colId] == '*' && len(nums) == 2 {
				ans += nums[0] * nums[1]
			}
		}
	}
	fmt.Println("ANS:", ans)
}

func isDigit(input rune) bool {
	return '0' <= input && input <= '9'
}

func isSymbol(rowId, colId int, digit string, input rune) bool {
	if input == '*' {
		num, _ := strconv.Atoi(digit)
		gears[rowId][colId] = append(gears[rowId][colId], num)
	}
	return (!isDigit(input) && input != '.')
}

func isPart(startId, endId int, num string, rowId int, matrix []string) bool {
	if startId > 0 {
		startId--
		if isSymbol(rowId, startId, num, rune(matrix[rowId][startId])) {
			return true
		}
	}
	if endId < len(matrix[rowId])-1 {
		endId++
		if isSymbol(rowId, endId, num, rune(matrix[rowId][endId])) {
			return true
		}
	}
	for start := startId; start <= endId; start++ {
		// Up
		if rowId > 0 && isSymbol(rowId-1, start, num, rune(matrix[rowId-1][start])) {
			return true
		}
		// Down
		if rowId < len(matrix)-1 && isSymbol(rowId+1, start, num, rune(matrix[rowId+1][start])) {
			return true
		}
	}
	return false
}
