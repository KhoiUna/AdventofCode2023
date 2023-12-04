package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ans := 0

	// Loop each game
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, ": ")[1]
		line = strings.ReplaceAll(line, ", ", ";")
		line = strings.ReplaceAll(line, "; ", ";")

		minCubes := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		subsets := strings.Split(line, ";")
		for _, subset := range subsets {
			numByte := strings.Split(subset, " ")[0]
			num, _ := strconv.Atoi(string(numByte))
			color := strings.Split(subset, " ")[1]
			if num > minCubes[color] {
				minCubes[color] = num
			}
		}

		power := 1
		for _, value := range minCubes {
			power *= value
		}
		ans += power
	}

	fmt.Println("ANS:", ans)
}
