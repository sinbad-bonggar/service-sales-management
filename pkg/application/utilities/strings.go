package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

// ArrayIntToString convert array integer to string array
// Example: ArrayIntToString([]int{1,2,3,4}, ",") -> 1,2,3,4
func ArrayIntToString(val []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(val)), delim), "[]")
}

// StringToArrayInt ...
//  Convert string delimited to array of integer
//  "1,2,3" -> []int{1, 2, 3}
func StringToArrayInt(str, delim string) ([]int, error) {
	sStrings := strings.Split(str, delim)

	result := []int{}
	for _, s := range sStrings {
		val, err := strconv.Atoi(s)
		if err != nil {
			return result, err
		}

		result = append(result, val)
	}

	return result, nil
}
