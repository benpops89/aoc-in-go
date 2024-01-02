package utils

import "strconv"

func StrToIntArray(s []string) ([]int, error) {
	var array []int
	for _, value := range s {
		i, err := strconv.Atoi(value)
		if err != nil {
			return []int{}, err
		}
		array = append(array, i)
	}

	return array, nil
}
