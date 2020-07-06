package numbers

import (
	"math"
)


func IsNarcissistic(number int) bool {

	var remainder float64
	var temp int
	var result float64 = 0.0
	var length float64 = float64(Digits(number))
	temp = number


	for {
		remainder = float64(temp % 10)
		result += math.Pow(remainder, length)
		temp /= 10

		if temp == 0.0 {
			break
		}
	}

	if number == int(result) {
		return true
	} else {
		return false
	}
}
