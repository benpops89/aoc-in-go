package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Card struct {
	game   int
	win    int
	copies int
}

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var cards []Card
	p1 := 0.0
	p2 := 0
	row := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		winning := []string{}

		card := strings.Split(scanner.Text(), ":")[1]
		numbers := strings.Split(strings.Trim(card, " "), "|")

		for _, w := range strings.Split(strings.Replace(strings.Trim(numbers[0], " "), "  ", " ", -1), " ") {
			for _, n := range strings.Split(strings.Replace(strings.Trim(numbers[1], " "), "  ", " ", -1), " ") {
				if w == n {
					winning = append(winning, n)
				}
			}
		}

		cards = append(cards, Card{row, len(winning), 1})
		row++

		if len(winning) > 0 {
			p1 += math.Pow(2.0, float64(len(winning)-1))
		}
	}

	fmt.Println(p1)

	for i, card := range cards {
		for j := i + 1; j <= i+card.win; j++ {
			cards[j].copies += cards[i].copies
		}
		p2 += cards[i].copies
	}

	fmt.Println(p2)
}
