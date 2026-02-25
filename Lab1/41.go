package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №41
// Напишіть код, який дозволяє користувачу вводити map[int]float32
// (подається як 2 списку) і число A. Видаліть всі елементи, значення ключів
// яких <= A. Виведіть в термінал отриманий результат і суму всіх значень Map.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter keys (comma-separated integers): ")
	keysInput, _ := reader.ReadString('\n')
	keyStrings := strings.Split(strings.TrimSpace(keysInput), ",")

	fmt.Print("Enter values (comma-separated floats): ")
	valuesInput, _ := reader.ReadString('\n')
	valueStrings := strings.Split(strings.TrimSpace(valuesInput), ",")

	mapData := make(map[int]float32)
	for i := 0; i < len(keyStrings) && i < len(valueStrings); i++ {
		key, _ := strconv.Atoi(strings.TrimSpace(keyStrings[i]))
		value, _ := strconv.ParseFloat(strings.TrimSpace(valueStrings[i]), 32)
		mapData[key] = float32(value)
	}

	fmt.Print("Enter threshold (A): ")
	aInput, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(strings.TrimSpace(aInput))

	for key := range mapData {
		if key <= a {
			delete(mapData, key)
		}
	}

	sum := float32(0)
	for _, value := range mapData {
		sum += value
	}

	fmt.Printf("Result: %v\n", mapData)
	fmt.Printf("Sum of all values: %f\n", sum)
}
