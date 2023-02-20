package euclid

import (
	"testing"
)

func TestAlgo(t *testing.T) {
	tests := []struct {
		m        int
		n        int
		expected int
	}{
		{
			m:        220,
			n:        40,
			expected: 20,
		},
		{
			m:        599,
			n:        15,
			expected: 1,
		},
		{
			m:        632,
			n:        32,
			expected: 8,
		},
	}

	for i, test := range tests {
		got := Algo(test.m, test.n)
		if got != test.expected {
			t.Errorf("test %v failed got = %v, expected = %v  \n", i, got, test.expected)
			t.Fail()
		}
	}
	t.Log("all tests passed")
}
