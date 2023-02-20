package euclid

// AlgorithmE given two positive integers m and n, find
// their greatest common divisor, that is, the largest positive integer that evenly
// divides both m and n.
func AlgorithmE(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	r := m % n
	if r == 0 {
		return n
	}

	return AlgorithmE(n, r)
}
