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

type SeedRange struct {
	start, end int
}

type Ranges = []Range
type Maps = []Ranges

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

	var seeds []SeedRange
	for _, seed := range StrToNumSlice(data[0][6:]) {
		seeds = append(seeds, SeedRange{seed, seed})
	}

	var part = 0
	var maps Maps = Maps{Ranges{}}

	for i := 3; i < len(data); i++ {
		if data[i] == "" {
			// End of section update
			part++
			maps = append(maps, Ranges{})
			i++
		} else {
			nums := StrToNumSlice(data[i])
			maps[part] = append(maps[part], Range{nums[0], nums[1], nums[2]})
		}
	}

	seed_map := make(map[SeedRange]SeedRange)

	for _, seed := range seeds {
		seed_map[seed] = seed
		for _, ranges := range maps {
			current := seed_map[seed]
			for _, r := range ranges {
				if current.start >= r.src && current.start < r.src+r.rng {
					start := r.dst + current.start - r.src
					seed_map[seed] = SeedRange{start, start}
					break
				}
			}
		}
	}

	p1 := math.MaxInt
	for _, s := range seed_map {
		p1 = min(p1, s.start)
	}

	// seed_source := make(map[SeedRange]SeedRange)
	//
	// for i := 0; i < len(seeds); i += 2 {
	// 	seed := SeedRange{seeds[i], seeds[i] + seeds[i+1]}
	// 	seed_source[seed] = seed
	// 	// var seedRange = Range{0, seeds[i], seeds[i+1]}
	// 	// var sourceRange = []Range{seedRange}
	// 	for _, ranges := range steps {
	// 		current := seed_source[seed]
	// 		for _, r := range ranges {
	// 			start := max(r.src, current.start)
	// 			end := min(r.src+r.rng, current.end)
	// 			offset := r.dst - r.src
	// 			if end > start {
	// 				seed_source[seed] = SeedRange{seed.start + offset, seed.end + offset}
	// 			}
	// 		}
	// 	}
	// }
	//
	// fmt.Println(seed_source)
	fmt.Println(seed_map)
	fmt.Println(p1)
}
