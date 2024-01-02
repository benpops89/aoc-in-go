package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Seed struct {
	start, end int
}

type Seeds []Seed

type Range struct {
	dst, src, length int
}

func main() {
	filename := os.Args[1]

	f, _ := os.ReadFile(filename)
	data := strings.Split(strings.TrimSpace(string(f)), "\n\n")

	numbers := strToNum(data[0][6:])

	var p1 []Seed
	var p2 []Seed
	for i := 0; i < len(numbers); i += 2 {
		p1 = append(p1, []Seed{{numbers[i], numbers[i] + 1}, {numbers[i+1], numbers[i+1] + 1}}...)
		p2 = append(p2, Seed{numbers[i], numbers[i] + numbers[i+1]})
	}

	parts := [][]Seed{p1, p2}
	for _, part := range parts {
		locations := GetLocations(data, part)

		sort.Slice(locations, func(i, j int) bool {
			return locations[i].start < locations[j].start
		})

		fmt.Println(locations[0].start)
	}
}

func GetLocations(s []string, seeds []Seed) []Seed {
	var seed Seed

	for _, section := range s[1:] {
		var ranges []Range

		for _, line := range strings.Split(section, "\n")[1:] {
			v := strToNum(line)
			ranges = append(ranges, Range{v[0], v[1], v[2]})
		}

		var locations []Seed
		var check bool
		for len(seeds) > 0 {
			seed, seeds = seeds[0], seeds[1:]
			for _, rng := range ranges {
				check = false
				start := max(seed.start, rng.src)
				end := min(seed.end, rng.src+rng.length)

				if start < end {
					locations = append(locations, Seed{start - rng.src + rng.dst, end - rng.src + rng.dst})

					if start > seed.start {
						seeds = append(seeds, Seed{seed.start, start})
					}

					if seed.end > end {
						seeds = append(seeds, Seed{end, seed.end})
					}
					check = true
					break
				}
			}
			if !check {
				locations = append(locations, seed)
			}
		}
		seeds = locations
	}

	return seeds
}

func strToNum(s string) []int {
	var numbers []int
	for _, i := range strings.Fields(strings.TrimSpace(s)) {
		n, _ := strconv.Atoi(i)
		numbers = append(numbers, n)
	}

	return numbers
}
