package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Завдання №57 (Вариант 15)
// На вхід подається ціле число n.
// Обчислити: tan(n) - 2n / sqrt(10 + 0.6n)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter n: ")
	nInput, _ := reader.ReadString('\n')
	n, _ := strconv.ParseFloat(strings.TrimSpace(nInput), 64)

	numerator := math.Tan(n) - 2*n
	denominator := math.Sqrt(10 + 0.6*n)

	if denominator == 0 {
		fmt.Println("Error: Division by zero!")
		return
	}

	result := numerator / denominator

	fmt.Printf("Result: %f\n", result)
}
