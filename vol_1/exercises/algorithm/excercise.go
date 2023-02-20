package algorithm

// Replacement [10] The text showed how to interchange the values of variables m and n,
// using the replacement notation, by setting t ← m, m ← n, n ← t. Show how the
// values of four variables (a, b, c, d) can be rearranged to (b, c, d, a) by a
// sequence of replacements. In other words, the new value of a is to be the
// original value of b, etc. Try to use the minimum number of replacements.
func Replacement(a, b, c, d int) (int, int, int, int) {
	a, b, c, d = b, c, d, a
	return a, b, c, d
}

//[15] Prove that m is always greater than n at the beginning of step E1,
//except possibly the first time this step occurs
//		ANSWER
//We get reminder m % n and replace n with this reminder and m with n, so
//m always will be greater than n

// AlgorithmF [20] Change Algorithm E (for the sake of efficiency) so that all trivial
// replacement operations such as “m ← n” are avoided. Write this new algorithm
// in the style of Algorithm E, and call it Algorithm F.
func AlgorithmF(m, n int) int {
	if m <= 0 || n <= 0 {
		return m
	}

	return AlgorithmF(n, m%n)
}

//[16] What is the greatest common divisor of 2166 and 6099?
//		ANSWER
// 57

//[12] Show that the “Procedure for Reading This Set of Books” that appears
//after the preface actually fails to be a genuine algorithm on at least three of our
//five counts! Also mention some differences in format between it and Algorithm E.
//		ANSWER
// no answer

//[20] What is T5, the average number of times step E1 is performed when n = 5?
