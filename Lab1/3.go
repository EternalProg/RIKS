package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Завдання №3
// Напишіть код, де користувач вводить рядок і букву.
// Порахуйте скільки разів задана буква входить в рядок і виведіть
// отриманий результат в термінал, а також індекс першого входження букви в рядок.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write string: ")
	strInput, _ := reader.ReadString('\n')
	str := strings.TrimSpace(strInput)

	fmt.Print("Write letter: ")
	letterInput, _ := reader.ReadString('\n')
	letterInput = strings.TrimSpace(letterInput)
	if len(letterInput) == 0 {
		fmt.Println("No letter entered!")
		return
	}
	letter := string([]rune(letterInput)[0])

	count := 0
	firstIndex := -1
	for i, ch := range str {
		if string(ch) == letter {
			count++
			if firstIndex == -1 {
				firstIndex = i
			}
		}
	}

	fmt.Printf("Letter '%s' appears %d times\n", letter, count)
	if firstIndex != -1 {
		fmt.Printf("First occurrence at index %d\n", firstIndex)
	} else {
		fmt.Printf("Letter '%s' not found in the string\n", letter)
	}
}
