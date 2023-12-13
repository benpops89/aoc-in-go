package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type Point struct {
	x int
	y int
}

type Number struct {
	number int
	row    int
	start  int
	end    int
}

type Grid struct {
	points  map[Point]rune
	numbers map[Number]bool
}

func (g *Grid) GetNumbers(pt Point) []int {
	var numbers []int
	seen := make(map[Point]bool)

	for _, p := range g.IsValid(pt) {
		seen[p] = true
		_, ok := seen[Point{p.x, p.y - 1}]
		for n, _ := range g.numbers {
			if n.row == p.x && p.y >= n.start && p.y < n.end && !ok {
				numbers = append(numbers, n.number)
			}
		}
	}

	return numbers
}

func (g *Grid) IsValid(p Point) []Point {
	var valid []Point

	points := []Point{
		Point{-1, -1}, Point{-1, 0}, Point{-1, 1},
		Point{0, -1}, Point{0, 1},
		Point{1, -1}, Point{1, 0}, Point{1, 1},
	}

	for _, pt := range points {
		check := Point{p.x + pt.x, p.y + pt.y}
		if v, ok := g.points[check]; ok && unicode.IsDigit(v) {
			valid = append(valid, check)
		}
	}

	return valid
}

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	points := make(map[Point]rune)
	parts := make(map[Number]bool)

	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		re := regexp.MustCompile(`\d+`)
		line := scanner.Text()
		numbers := re.FindAllStringSubmatchIndex(line, -1)

		for _, n := range numbers {
			value, _ := strconv.Atoi(line[n[0]:n[1]])
			parts[Number{value, row, n[0], n[1]}] = false
		}

		for col, char := range line {
			if unicode.IsDigit(char) || string(char) != "." {
				points[Point{row, col}] = char
			}
		}

		row++
	}

	grid := Grid{points, parts}

	p1 := 0
	p2 := 0

	for pt, value := range grid.points {
		if !unicode.IsDigit(value) {
			for _, number := range grid.GetNumbers(pt) {
				p1 += number
			}
		}

		if string(value) == "*" {
			if numbers := grid.GetNumbers(pt); len(numbers) > 1 {
				p2 += numbers[0] * numbers[1]
			}
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
