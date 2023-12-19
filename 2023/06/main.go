package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]

	f, _ := os.ReadFile(filename)
	data := strings.Split(strings.TrimSpace(string(f)), "\n")

	var races [][]int
	for _, lines := range data {
		numbers := strToNum(strings.Split(lines, ":")[1])
		races = append(races, numbers)
	}

	// fmt.Println(races)

	p1 := 1
	for i := 0; i < len(races[0]); i++ {
		var distances []int
		for j := 1; j < races[0][i]; j++ {
			distance := j * (races[0][i] - j)
			if distance > races[1][i] {
				distances = append(distances, distance)
			}
		}
		p1 *= len(distances)
	}
	fmt.Println(p1)
}

func strToNum(s string) []int {
	var numbers []int
	for _, i := range strings.Fields(strings.TrimSpace(s)) {
		n, _ := strconv.Atoi(i)
		numbers = append(numbers, n)
	}

	return numbers
}
