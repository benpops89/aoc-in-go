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
	X int
	Y int
}

func (pt *Point) AddValue(p Point) Point {
	return Point{pt.X + p.X, pt.Y + p.Y}
}

func CreateAdjacent() []Point {
	points := []Point{}
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			points = append(points, Point{i, j})
		}
	}

	return points
}

func FindNumbers(r string) [][]int {
	var result [][]int
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringSubmatchIndex(r, -1)
	for _, match := range matches {
		value := r[match[0]:match[1]]
		number, _ := strconv.Atoi(value)
		result = append(result, append(match, number))
	}

	return result
}

type Part struct {
	row   int
	start int
	end   int
}

func main() {
	filename := os.Args[1]
	scanner := OpenFile(filename)

	engine := make(map[Point]rune)
	numbers := make(map[int][][]int)
	var parts [][]int

	line := 0
	for scanner.Scan() {
		numbers[line] = FindNumbers(scanner.Text())
		for i, char := range scanner.Text() {
			if unicode.IsDigit(char) || char == 42 {
				engine[Point{line, i}] = char
			}
		}

		line++
	}
	fmt.Println(engine)
	fmt.Println(len(engine))
	fmt.Println(numbers)

	for coord, value := range engine {
		if value == 42 {
			neighbours := CheckNeighbours(coord, engine)
			if len(neighbours) > 0 {
				// Find neighbouring numbers
				var part []int
				for _, n := range neighbours {
					for _, d := range numbers[n.X] {
						if coord.Y >= d[0]-1 && n.Y <= d[1] {
							part = append(part, d[2])
						}
					}
				}
				parts = append(parts, part)
				fmt.Println(neighbours)
			}

		}
	}
	fmt.Println(parts)
	sum := 0
	for _, values := range parts {
		seen := make(map[int]bool)
		mult := 1
		for _, value := range values {
			if _, ok := seen[value]; !ok {
				seen[value] = true
				mult *= value
			}
		}
		if len(seen) > 1 {
			sum += mult
		}

	}
	fmt.Println(sum)
}

func CheckNeighbours(c Point, engine map[Point]rune) []Point {
	var neighbours []Point
	for _, point := range CreateAdjacent() {
		p := c.AddValue(point)
		n, ok := engine[c.AddValue(point)]
		if ok && unicode.IsDigit(n) {
			neighbours = append(neighbours, p)
		}
	}
	return neighbours
}

func OpenFile(s string) *bufio.Scanner {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(f)
	return scanner
}
