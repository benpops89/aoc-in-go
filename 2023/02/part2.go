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

		power := 1
		for _, s := range values {
			power *= s
		}

		sum += power
	}
	fmt.Println(sum)
}
