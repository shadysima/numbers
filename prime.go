package numbers

import "math"


func IsPrime(num uint32) bool {
	
	var n int = int(num)
	if num == 2 {
		return true
	}
	if num % 2 == 0 {
		return false
	}
	
	var r int = int(math.Floor(math.Sqrt(float64(num))))
	var step int = 3

	for step < r {
		if n % step == 0 {
			return false
		}
		step += 2
	}

	return true
}
