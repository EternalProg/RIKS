package main

import "fmt"

/*
18.Користувач вводить ціле число, що представляє номер місяця. Напишіть
функцію, яка повертає кількість днів у цьому місяці. Якщо введено
некоректний номер місяця, функція повинна повертати нуль. Отриманий
результат виведіть у термінал.
*/

func task18(month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		return 28
	default:
		return 0
	}
}

func main() {
	var month int
	fmt.Scan(&month)

	fmt.Println(task18(month))
}
