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
	ids := []string{}

	// Loop each game
	for scanner.Scan() {
		failed := false
		line := scanner.Text()
		game := strings.Split(line, ": ")[1]

		id := strings.Split(strings.ReplaceAll(line, " ", ""), ":")[0][4:]
		sets := strings.Split(game, ";")

		// Loop each set
		for _, set := range sets { // set: 1 red,2 green,6 blue
			cubes := map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			}
			subsets := strings.Split(set, ",")

			// Loop each subset
			for _, subset := range subsets { // subset: 1 red
				trimmed := strings.TrimSpace(subset)
				numByte := strings.Split(trimmed, " ")[0]
				num, _ := strconv.Atoi(string(numByte))
				color := strings.Split(trimmed, " ")[1]

				cubes[color] -= num

				if cubes[color] < 0 {
					failed = true
					break
				}
			}

			if failed {
				break
			}
		}

		if failed {
			continue
		}

		ids = append(ids, id)
	}

	ans := 0
	for _, id := range ids {
		num, _ := strconv.Atoi(id)
		ans += num
	}

	fmt.Println("ANS:", ans)
}
