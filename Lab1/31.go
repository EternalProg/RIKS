package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №31
// Напишіть код, який дозволяє користувачу вводити зріс цілих чисел і два
// числа (наприклад a і b). Порахуйте суму елементів зрізу, які лежать в
// діапазоні від індексу «a» по «b». Виведіть в термінал отриманий результат.

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

	fmt.Print("Enter start index (a): ")
	aInput, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(strings.TrimSpace(aInput))

	fmt.Print("Enter end index (b): ")
	bInput, _ := reader.ReadString('\n')
	b, _ := strconv.Atoi(strings.TrimSpace(bInput))

	if a < 0 || b >= len(numbers) || a > b {
		fmt.Println("Invalid range!")
		return
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += numbers[i]
	}

	fmt.Printf("Sum of elements from index %d to %d: %d\n", a, b, sum)
}
