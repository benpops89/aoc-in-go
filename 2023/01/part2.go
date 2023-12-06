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

	// Create a map from name to number
	number_map := make(map[string]string)
	number_map["one"] = "1"
	number_map["two"] = "2"
	number_map["three"] = "3"
	number_map["four"] = "4"
	number_map["five"] = "5"
	number_map["six"] = "6"
	number_map["seven"] = "7"
	number_map["eight"] = "8"
	number_map["nine"] = "9"

	sum := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		re_first := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|[1-9]`)
		first_number := re_first.FindString(scanner.Text())

		first_value, ok := number_map[first_number]
		if ok {
			first_number = first_value
		}

		re_last := regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|[1-9])`)
		last_number := re_last.FindStringSubmatch(scanner.Text())[1]

		last_value, ok := number_map[last_number]
		if ok {
			last_number = last_value
		}

		new_value, err := strconv.Atoi(first_number + last_number)
		if err != nil {
			log.Fatal(err)
		}
		sum += new_value
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading data:", err)
	}

	fmt.Println(sum)
}
