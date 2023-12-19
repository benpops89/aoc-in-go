package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	f, _ := os.ReadFile(filename)

	data := strings.Split(string(f), "\n")

	fmt.Println(data)
}
