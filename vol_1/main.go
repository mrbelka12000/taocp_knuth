package main

import (
	"fmt"
	"github.com/mrbelka12000/taocp_knuth/vol_1/euclid"
	"github.com/mrbelka12000/taocp_knuth/vol_1/exercises/algorithm"
)

func main() {
	res := euclid.AlgorithmE(220, 40)
	fmt.Println(res)

	res = algorithm.AlgorithmF(220, 40)
	fmt.Println(res)

	fmt.Println(algorithm.Replacement(1, 2, 3, 4))
}
