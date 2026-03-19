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
		sum += ((nF*nF + 5) * 16 * math.Pow(3, nF)) / 25
	}

	fmt.Println(sum)
}
