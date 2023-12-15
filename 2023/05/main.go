package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	dst, src, rng int
}

type Ranges = []Range
type Steps = []Ranges

func Convert(start int, dst int, src int, rng int) int {
	fmt.Println("\t\t", start, dst, src, rng)
	if start >= src && start < src+rng {
		return dst + (start - src)
	}

	return start
}

func StrToNumSlice(s string) []int {
	var numbers []int
	for _, i := range strings.Fields(strings.TrimSpace(s)) {
		n, _ := strconv.Atoi(i)
		numbers = append(numbers, n)
	}

	return numbers
}

func main() {
	filename := os.Args[1]

	f, _ := os.ReadFile(filename)
	data := strings.Split(string(f), "\n")

	seeds := StrToNumSlice(data[0][6:])

	var step = 0
	var steps Steps = Steps{Ranges{}}

	for i := 3; i < len(data); i++ {
		if data[i] == "" {
			// End of section update
			step++
			steps = append(steps, Ranges{})
			i++
		} else {
			nums := StrToNumSlice(data[i])
			steps[step] = append(steps[step], Range{nums[0], nums[1], nums[2]})
		}
	}

	seed_map := make(map[int]int)

	for _, seed := range seeds {
		seed_map[seed] = seed
		for _, ranges := range steps {
			current := seed_map[seed]
			for _, r := range ranges {
				if current >= r.src && current < r.src+r.rng {
					seed_map[seed] = r.dst + current - r.src
					break
				}
			}
		}
	}

	p1 := math.MaxInt
	for _, s := range seed_map {
		p1 = min(p1, s)
	}

	fmt.Println(seed_map)
	fmt.Println(p1)
}
