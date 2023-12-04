package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ans := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digits := []string{}

		for _, r := range line {
			if unicode.IsDigit(r) {
				digits = append(digits, string(r))
			}
		}
		num, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
		ans += num
	}
	fmt.Println(ans)
}
