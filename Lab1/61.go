package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №61 (Вариант 19)
// На вхід подається ціле число n.
// Обчислити: (2*n^2 - 4*n + 10) / 2*n

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter n: ")
	nInput, _ := reader.ReadString('\n')
	n, _ := strconv.ParseFloat(strings.TrimSpace(nInput), 64)

	if n == 0 {
		fmt.Println("Error: Division by zero!")
		return
	}

	numerator := 2*n*n - 4*n + 10
	denominator := 2 * n

	result := numerator / denominator

	fmt.Printf("Result: %f\n", result)
}
