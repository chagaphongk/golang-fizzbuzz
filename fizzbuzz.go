package fizzbuzz

import "strconv"

// Say n
func Say(n int) string {

	if n > 5 {
		return strconv.Itoa(n) + " is not support"
	}

	if n == 3 {
		return "fizz"
	} else if n == 5 {
		return "buzz"
	}

	return strconv.Itoa(n)
}
