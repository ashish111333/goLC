package stack

// Note: optimize it currently use string better use byte slice
func isValid(s string) bool {
	left_bracket := "("
	right_bracket := ")"
	left_brace := "{"
	right_brace := "}"
	left_sq_bracket := "["
	right_sq_bracket := "]"

	scanned_str := ""

	for _, v := range s {
		if string(v) == left_brace {
			scanned_str += left_brace

		}
		if string(v) == left_bracket {
			scanned_str += left_bracket
		}
		if string(v) == left_sq_bracket {
			scanned_str += left_sq_bracket
		}
		if string(v) == right_brace {
			if len(scanned_str) == 0 {
				return false
			}
			if string(scanned_str[len(scanned_str)-1]) != left_brace {
				return false
			} else {
				scanned_str = scanned_str[:len(scanned_str)-1]
			}

		}
		if string(v) == right_bracket {
			if len(scanned_str) == 0 {
				return false
			}
			if string(scanned_str[len(scanned_str)-1]) != left_bracket {
				return false
			} else {
				scanned_str = scanned_str[:len(scanned_str)-1]

			}
		}
		if string(v) == right_sq_bracket {
			if len(scanned_str) == 0 {
				return false
			}
			if string(scanned_str[len(scanned_str)-1]) != left_sq_bracket {
				return false
			} else {
				scanned_str = scanned_str[:len(scanned_str)-1]
			}

		}

	}
	if len(scanned_str) != 0 {
		return false
	}
	return true

}
