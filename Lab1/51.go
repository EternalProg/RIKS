package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №51
// Користувач вводить зріс mySlice з цілими числами.
// Знайти добуток його другого і середнього елементів і вивести отриманий результат в термінал.

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

	if len(numbers) < 2 {
		fmt.Println("Need at least 2 elements")
		return
	}

	second := numbers[1]
	middle := numbers[len(numbers)/2]
	product := second * middle

	fmt.Printf("Product of second (%d) and middle (%d) elements: %d\n", second, middle, product)
}
