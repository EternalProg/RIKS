package main

import "fmt"

func main() {
	var z int
	if _, err := fmt.Scan(&z); err != nil {
		return
	}

	sum := 0.0
	for n := 1; n <= z; n++ {
		nF := float64(n)
		sum += (2*nF*nF - 4*nF + 10) / (2 * nF)
	}

	fmt.Println(sum)
}
