package numbers

import (
	"strconv"
	"fmt"
)


func ToArray(num int) []int {

	var final []int
	var temp []string
	var s string = strconv.Itoa(num)

	for _, val := range s {
		temp = append(temp, fmt.Sprintf("%c", val))
	}

	for _, val := range temp {
		strint, _ := strconv.Atoi(val)
		final = append(final, strint)
	}

	return final

}
