package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	f, _ := os.ReadFile(filename)

	fmt.Println(string(f))
}
