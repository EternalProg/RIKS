package main

import "fmt"

/*
21.Користувач вводить довжину, ширину та висоту коробки. Напишіть
функцію, яка повертає об’єм коробки та має два аргументи за
замовчуванням (ширина = 10, висота = 7) на випадок, якщо користувач
введе не три значення. Отриманий результат виведіть у термінал.
*/

func task21(length int, width_height ...int) int {
	width := 10
	height := 7

	if len(width_height) > 0 {
		width = width_height[0]
	}
	if len(width_height) > 1 {
		height = width_height[1]
	}

	return length * width * height
}

func main() {
	var length, width, height int

	count, _ := fmt.Scan(&length, &width, &height)
	if count == 1 {
		fmt.Println(task21(length))
		return
	}
	if count == 2 {
		fmt.Println(task21(length, width))
		return
	}
	if count == 3 {
		fmt.Println(task21(length, width, height))
	}
}
