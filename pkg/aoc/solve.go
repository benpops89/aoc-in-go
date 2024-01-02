package aoc

import (
	"fmt"
	"os"
	"strings"
)

func Solve(filename string) {
	file, _ := os.ReadFile(filename)
	data := strings.Split(strings.TrimSpace(string(file)), "\n")
	fmt.Println(data)
	fmt.Println("Add additional information")
}
