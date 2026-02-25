package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Завдання №7
// Напишіть код, де користувач вводить два рядки str1 і str2.
// Програма повинна створити новий рядок str3 шляхом додавання str2 в середину str1.
// Отриманий результат виведіть в термінал.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write first string (str1): ")
	str1Input, _ := reader.ReadString('\n')
	str1 := strings.TrimSpace(str1Input)

	fmt.Print("Write second string (str2): ")
	str2Input, _ := reader.ReadString('\n')
	str2 := strings.TrimSpace(str2Input)

	middle := len(str1) / 2
	str3 := str1[:middle] + str2 + str1[middle:]

	fmt.Printf("Result (str3): %s\n", str3)
}
