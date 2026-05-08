package main

import (
	"fmt"
)

func main() {
	if err := task1FileCopyGoroutines("input.txt", "output.txt"); err != nil {
		fmt.Println("Завдання 1 завершилось з помилкою:", err)
		return
	}

	task3DiningPhilosophers()
	task7SumOfMeans()
}
