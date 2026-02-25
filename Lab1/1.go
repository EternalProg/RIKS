package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Напишіть код, де користувач вводить рядок і букву, наявність якої
// необхідно перевірити у введеному рядку. Виведіть в термінал отриманий
// результат, а також індекс останнього входження букви в рядок.

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

	lastIndex := -1
	for i := len(str) - 1; i >= 0; i-- {
		if string([]rune(str)[i]) == letter {
			lastIndex = i
			break
		}
	}

	if lastIndex != -1 {
		fmt.Printf("Letter '%s' found at index %d\n", letter, lastIndex)
	} else {
		fmt.Printf("Letter '%s' not found in the string\n", letter)
	}
}
