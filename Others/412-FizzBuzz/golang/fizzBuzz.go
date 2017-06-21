package fizzBuzz

import "strconv"

func fizzBuzz(n int) []string {
	var result []string
	for i, fizz, buzz := 1, 0, 0; i <= n; i++ {
		fizz++
		buzz++
		if fizz == 3 && buzz == 5 {
			result = append(result, "FizzBuzz")
			fizz, buzz = 0, 0
		} else if fizz == 3 {
			result = append(result, "Fizz")
			fizz = 0
		} else if buzz == 5 {
			result = append(result, "Buzz")
			buzz = 0
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
