package dp

import (
	"strconv"
)

func numDecodings(s string) int {

	var get_ways_of_decodings func(int) int

	// do we have leading zeroes...  ?
	first_digit, _ := strconv.Atoi(string(s[0]))

	if first_digit == 0 {
		return 0
	}
	if len(s) == 2 {
		intg, _ := strconv.Atoi(s)
		if string(s[1]) == "0" {
			if intg < 26 {
				return 1
			}
		} else {
			if intg <= 26 {
				return 2
			} else {
				return 1
			}
		}
	}
	isDoubleDigitIntValid := func(s string) bool {
		if len(s) != 2 {
			panic("invalid	 input")

		}
		if string(s[0]) == "0" {
			return false
		}

		intg, _ := strconv.Atoi(s)
		if intg <= 26 {
			return true
		}
		return false

	}

	computed_decodings := map[int]int{}

	get_ways_of_decodings = func(i int) int {

		if i == 0 {
			return 1
		}
		if v, ok := computed_decodings[i]; ok {
			return v
		}

		if string(s[i]) == "0" {

			if isDoubleDigitIntValid(s[i-1 : i+1]) {
				computed_decodings[i] = get_ways_of_decodings(i - 1)
			} else {
				return 0
			}

		} else {
			if isDoubleDigitIntValid(s[i-1 : i+1]) {
				computed_decodings[i] = get_ways_of_decodings(i-1) + 1
			} else {
				computed_decodings[i] = get_ways_of_decodings(i - 1)
			}

		}
		return computed_decodings[i]

	}
	return get_ways_of_decodings(len(s) - 1)

}
