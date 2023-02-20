package euclid

// Algo given two positive integers m and n, find
// their greatest common divisor, that is, the largest positive integer that evenly
// divides both m and n.
func Algo(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	r := m % n
	if r == 0 {
		return n
	}

	return Algo(n, r)
}
