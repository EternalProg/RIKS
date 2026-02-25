package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №39
// Напишіть код, який дозволяє користувачу вводити map[int]string
// (подається як 2 списки) і рядок А. Видаліть всі елементи, значення яких
// починаються з підрядка А і виведіть в термінал отриманий результат.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter keys (comma-separated integers): ")
	keysInput, _ := reader.ReadString('\n')
	keyStrings := strings.Split(strings.TrimSpace(keysInput), ",")

	fmt.Print("Enter values (comma-separated strings): ")
	valuesInput, _ := reader.ReadString('\n')
	valueStrings := strings.Split(strings.TrimSpace(valuesInput), ",")

	mapData := make(map[int]string)
	for i := 0; i < len(keyStrings) && i < len(valueStrings); i++ {
		key, _ := strconv.Atoi(strings.TrimSpace(keyStrings[i]))
		value := strings.TrimSpace(valueStrings[i])
		mapData[key] = value
	}

	fmt.Print("Enter substring (A): ")
	substring, _ := reader.ReadString('\n')
	substring = strings.TrimSpace(substring)

	for key, value := range mapData {
		if strings.HasPrefix(value, substring) {
			delete(mapData, key)
		}
	}

	fmt.Printf("Result: %v\n", mapData)
}
