package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №30
// Напишіть код, який дозволяє користувачу вводити зріс цілих чисел.
// Сформуйте новий зріс, елементи якого – непарні значення.
// Виведіть в термінал отриманий результат.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter integers (comma-separated): ")
	input, _ := reader.ReadString('\n')
	stringNumbers := strings.Split(strings.TrimSpace(input), ",")

	var numbers []int
	for _, str := range stringNumbers {
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		numbers = append(numbers, num)
	}

	var oddNumbers []int
	for _, num := range numbers {
		if num%2 != 0 {
			oddNumbers = append(oddNumbers, num)
		}
	}

	fmt.Printf("Odd numbers: %v\n", oddNumbers)
}
