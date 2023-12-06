package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	sum := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		re := regexp.MustCompile(`\d`)
		numbers := re.FindAllString(scanner.Text(), -1)

		value, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if err != nil {
			log.Fatal(err)
		}
		sum += value
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading data:", err)
	}

	fmt.Println(sum)
}
