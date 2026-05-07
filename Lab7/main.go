package main

import (
	"fmt"
	"os"
)

const (
	outputDir = "output"
)

func main() {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		fmt.Println("Помилка створення папки output:", err)
		return
	}

	if err := task1_2(); err != nil {
		fmt.Println("Завдання 1.2 завершилось з помилкою:", err)
		return
	}

	if err := task1_6(); err != nil {
		fmt.Println("Завдання 1.6 завершилось з помилкою:", err)
		return
	}

	if err := task1_8(); err != nil {
		fmt.Println("Завдання 1.8 завершилось з помилкою:", err)
		return
	}

	if err := task1_12(); err != nil {
		fmt.Println("Завдання 1.12 завершилось з помилкою:", err)
		return
	}

	if err := task2JSON(); err != nil {
		fmt.Println("Завдання 2 завершилось з помилкою:", err)
		return
	}

	fmt.Println("\nУсі завдання виконано успішно.")
}
