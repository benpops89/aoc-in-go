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
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	outcome := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		values := make(map[string]int)
		info := strings.Split(scanner.Text(), ":")

		for _, game := range strings.Split(info[1], ";") {
			colors := strings.Split(strings.Trim(game, " "), ", ")

			for _, color := range colors {
				value := strings.Split(color, " ")

				current, ok := values[value[1]]
				v, err := strconv.Atoi(value[0])
				if err != nil {
					log.Fatal(err)
				}

				if ok {
					v = max(current, v)
				}
				values[value[1]] = v
			}
		}

		valid := true
		for c, v := range values {
			if v > outcome[c] {
				valid = false
				break
			}
		}

		if valid {
			game_id, err := strconv.Atoi(strings.Split(info[0], " ")[1])
			if err != nil {
				log.Fatal(err)
			}
			sum += game_id
		}
	}
	fmt.Println(sum)
}
