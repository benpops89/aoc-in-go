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
	parts := make(map[Part]int)

	line := 0
	for scanner.Scan() {
		numbers[line] = FindNumbers(scanner.Text())
		for i, char := range scanner.Text() {
			if char != 46 {
				engine[Point{line, i}] = char
			}
		}

		line++
	}
	fmt.Println(engine)
	fmt.Println(len(engine))
	fmt.Println(numbers)

	for coord, value := range engine {
		if unicode.IsDigit(value) {
			valid := CheckNeighbours(coord, engine)
			if valid {
				// Find neighbouring numbers
				fmt.Println(coord, " -- ", value)
				for _, n := range numbers[coord.X] {
					if coord.Y >= n[0]-1 && coord.Y <= n[1] {
						parts[Part{coord.X, n[0], n[1]}] = n[2]
					}
				}
			}

		}
	}
	fmt.Println(parts)
	sum := 0
	for _, v := range parts {
		sum += v
	}
	fmt.Println(sum)
}

func CheckNeighbours(c Point, engine map[Point]rune) bool {
	for _, point := range CreateAdjacent() {
		n, ok := engine[c.AddValue(point)]
		if ok && !unicode.IsDigit(n) {
			return true
		}
	}
	return false
}

func OpenFile(s string) *bufio.Scanner {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(f)
	return scanner
}
