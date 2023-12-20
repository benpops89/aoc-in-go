package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	symbol string
	value  int
}

type Hand struct {
	cards  []Card
	counts map[string]int
	score  []int
	bid    int
}

func main() {
	filename := os.Args[1]
	f, _ := os.ReadFile(filename)

	data := strings.Split(strings.TrimSpace(string(f)), "\n")

	strength := make(map[string]int)
	cards := []string{"A", "K", "Q", "J", "T"}
	i := 14
	for _, card := range cards {
		strength[card] = i
		i--
	}

	var hands []Hand
	for _, line := range data {
		var cards []Card
		s := strings.Split(line, " ")
		counts := make(map[string]int)
		for _, card := range s[0] {
			c := string(card)
			value, err := strconv.Atoi(c)
			if err != nil {
				value = strength[c]
			}
			cards = append(cards, Card{c, value})

			// Update counts
			if _, ok := counts[c]; ok {
				counts[c] += 1
			} else {
				counts[c] = 1
			}
		}
		bid, _ := strconv.Atoi(s[1])
		score := CalculateScore(counts)
		hand := Hand{cards, counts, score, bid}
		hands = append(hands, hand)
	}
	// Sort hands
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score[1] > hands[j].score[1] {
			return true
		} else if hands[i].score[1] == hands[j].score[1] {
			if hands[i].score[0] > hands[j].score[0] {
				return true
			} else if hands[i].score[0] == hands[j].score[0] {
				for k := 0; k < len(hands[i].cards); k++ {
					if hands[i].cards[k].value > hands[j].cards[k].value {
						return true
					} else if hands[i].cards[k].value < hands[j].cards[k].value {
						return false
					}

				}
			}
		}
		return false
	})
	p1 := 0
	rank := len(hands)
	for _, hand := range hands {
		p1 += rank * hand.bid
		rank--
	}

	fmt.Println(p1)

}

func CalculateScore(c map[string]int) []int {
	var scores []int
	for _, value := range c {
		scores = append(scores, value)
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })

	length := len(scores)
	if length == 1 {
		return []int{scores[0], scores[0]}
	}

	return []int{scores[length-2], scores[length-1]}
}
