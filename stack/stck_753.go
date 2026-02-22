package stack

// scan asteroids from left to right
// worst case TC O(2n)
func asteroidCollision(asteroids []int) []int {

	mod := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	scanned_asteroids := []int{}

	for _, v := range asteroids {

		if len(scanned_asteroids) == 0 {
			scanned_asteroids = append(scanned_asteroids, v)
			continue
		}
		// check when collision happens only when v is negative
		// then check who survives

		if v < 0 {

			start := len(scanned_asteroids) - 1
			for start >= 0 {
				if scanned_asteroids[start] > 0 {
					if mod(v) > scanned_asteroids[start] {
						scanned_asteroids = scanned_asteroids[:start]
						start--
						continue
					}
					if mod(v) == scanned_asteroids[start] {
						scanned_asteroids = scanned_asteroids[:start]
						start--
						if start < 0 {
							start++
						}
						break
					}
					if mod(v) < scanned_asteroids[start] {
						break
					}
				}

				start--

			}

			if start < 0 {
				scanned_asteroids = append(scanned_asteroids, v)

			}

			continue
		}
		scanned_asteroids = append(scanned_asteroids, v)

	}
	return scanned_asteroids

}
