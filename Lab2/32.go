package main

import (
	"fmt"
	"math"
)

func main() {
	var z int
	if _, err := fmt.Scan(&z); err != nil {
		return
	}

	sum := 0.0
	for n := 1; n <= z; n++ {
		nF := float64(n)
		sum += (nF - 20) / math.Sqrt(nF*nF*nF)
	}

	fmt.Println(sum)
}
