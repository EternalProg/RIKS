package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Завдання №58 (Вариант 16)
// На вхід подається ціле число n.
// Обчислити: (3*sin(n) - 15) / sqrt(n^5)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter n: ")
	nInput, _ := reader.ReadString('\n')
	n, _ := strconv.ParseFloat(strings.TrimSpace(nInput), 64)

	if n == 0 {
		fmt.Println("Error: n cannot be 0!")
		return
	}

	numerator := 3*math.Sin(n) - 15
	denominator := math.Sqrt(math.Pow(n, 5))

	result := numerator / denominator

	fmt.Printf("Result: %f\n", result)
}
