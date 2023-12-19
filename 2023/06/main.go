package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Races [][]int

func main() {
	filename := os.Args[1]

	f, _ := os.ReadFile(filename)
	data := strings.Split(strings.TrimSpace(string(f)), "\n")

	var one [][]int
	var two [][]int

	for _, line := range data {
		// Part 1
		one = append(one, strToNum(strings.Split(line, ":")[1]))

		// Part 2
		numbers := strToNum(strings.ReplaceAll(strings.Split(line, ":")[1], " ", ""))
		two = append(two, numbers)
	}

	p1 := CalculateDistances(one)
	fmt.Println(p1)

	p2 := CalculateDistances(two)
	fmt.Println(p2)
}

func strToNum(s string) []int {
	var numbers []int
	for _, i := range strings.Fields(strings.TrimSpace(s)) {
		n, _ := strconv.Atoi(i)
		numbers = append(numbers, n)
	}

	return numbers
}

func CalculateDistances(r [][]int) int {
	total := 1
	for i := 0; i < len(r[0]); i++ {
		var distances []int
		for j := 1; j < r[0][i]; j++ {
			distance := j * (r[0][i] - j)
			if distance > r[1][i] {
				distances = append(distances, distance)
			}
		}
		total *= len(distances)
	}
	return total
}
