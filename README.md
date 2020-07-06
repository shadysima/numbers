# numbers
A Go package that contains number utilities

## Function List

* `IsPrime(n uint32)`: Takes `n` and returns `bool` if `n` is prime.
* `IsNarcissistic(n int)`: Takes `n` and returns `bool` if `n` is a narcissistic number (Armstrong number)
* `IsDivisble(n1 int64, n2 int64)`: Takes `n1` and `n2` and returns `bool` if `n1` is evenly divisble by `n2`. Synonymous with `if n1 % n2 == 0`
* `Digits(n int)`: Takes `n` and return the number of digits as `int`
* `ToArray(n int)`: Takes `n` and returns a []int array whose elements are `n`'s individual digits.
	**EX** `numbers.ToArray(12345)` returns `[1 2 3 4 5]`
