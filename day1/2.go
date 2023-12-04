package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := 0
	for scanner.Scan() {
		digits := []int{}
		line := scanner.Text()

		for i, r := range line {
			if unicode.IsDigit(r) {
				d, _ := strconv.Atoi(string(r))
				digits = append(digits, d)
			}

			for d, val := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], val) {
					digits = append(digits, d+1)
				}
			}
		}
		num, _ := strconv.Atoi(strconv.Itoa(digits[0]) + strconv.Itoa(digits[len(digits)-1]))
		ans += num
	}

	fmt.Println("ANS:", ans)
}
