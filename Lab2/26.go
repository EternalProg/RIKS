package main

import "fmt"

// Користувач вводить зріз цілих чисел. Використовуючи цикл порахуйте кількість входжень в нього елементів кратних 5 і виведіть в термінал отриманий результат.

func main() {
	count := 0
	for {
		var v int
		if _, err := fmt.Scan(&v); err != nil {
			break
		}
		if v%5 == 0 {
			count++
		}
	}

	fmt.Println(count)
}
