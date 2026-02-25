package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №28
// Напишіть код, який дозволяє користувачу вводити зріз рядкових значень і
// два числа (наприклад a і b). Код повинен видалити із зрізу всі елементи,
// які лежать в діапазоні від індексу «a» по «b». Виведіть в термінал отриманий результат.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter strings (comma-separated): ")
	stringsInput, _ := reader.ReadString('\n')
	stringsList := strings.Split(strings.TrimSpace(stringsInput), ",")

	fmt.Print("Enter start index (a): ")
	aInput, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(strings.TrimSpace(aInput))

	fmt.Print("Enter end index (b): ")
	bInput, _ := reader.ReadString('\n')
	b, _ := strconv.Atoi(strings.TrimSpace(bInput))

	if a < 0 || b >= len(stringsList) || a > b {
		fmt.Println("Invalid range!")
		return
	}

	result := append(stringsList[:a], stringsList[b+1:]...)

	fmt.Printf("Result: %v\n", result)
}
