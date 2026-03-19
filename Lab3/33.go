package main

import "fmt"

/*
33.Користувач вводить рядок str і символ symbol. Напишіть рекурсивну
функцію, яка видаляє з рядка всі букви, які відповідають symbol.
Отриманий результат виведіть в термінал.
*/

func task33(str string, symbol rune) string {
	if str[0] == symbol {
		return task33(str[1:], symbol)
	}

	return string(str[0]) + task33(str[1:], symbol)
}

func main() {
	var str string
	var symbol rune
	fmt.Scan(&str, &symbol)

	fmt.Println(task33(str, sym))
}
