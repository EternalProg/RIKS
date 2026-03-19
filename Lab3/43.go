package main

import "fmt"

/*
43.Напишіть функцію на вхід якої подається рядок. Вона повинна повертати
іншу функцію, яка приймає номер індексу та повертає символ, який
розміщений в рядку по даному індексу. Якщо індекс виходить за межі
рядка, тоді повертається пустий символ. Отриманий результат виведіть в
термінал.
*/

func task43(s string) func(int) string {
	return func(index int) string {
		if index < 0 || index >= len(s) {
			return ""
		}
		return string(s[index])
	}
}

func main() {
	var s string
	var index int
	fmt.Scan(&s, &index)

	getter := task43(s)
	fmt.Println(getter(index))
}
