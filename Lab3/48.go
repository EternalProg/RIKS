package main

import "fmt"

/*
48.Напишіть функцію на вхід якої подається номер біта, який перевіряється.
Функція повинна повертати іншу функцію, яка приймає число і повертає
true, якщо заданий біт в ньому встановлений як одиниця, в іншому випадку
– false. Отриманий результат виведіть в термінал.
*/

func task48(bit int) func(int) bool {
	return func(value int) bool {
		return value&(1<<bit) != 0
	}
}

func main() {
	var bit, value int
	fmt.Scan(&bit, &value)

	check := task48(bit)
	fmt.Println(check(value))
}
