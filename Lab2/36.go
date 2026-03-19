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
		sum += (3*math.Sin(nF) - 15) / math.Sqrt(math.Pow(nF, 5))
	}

	fmt.Println(sum)
}
