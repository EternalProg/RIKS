package main

import "fmt"

/*
33.Користувач вводить рядок str і символ symbol. Напишіть рекурсивну
функцію, яка видаляє з рядка всі букви, які відповідають symbol.
Отриманий результат виведіть в термінал.
*/

func task33Bytes(s string, sym byte) string {
	if len(s) == 0 {
		return ""
	}
	if s[0] == sym {
		return task33Bytes(s[1:], sym)
	}
	return string(s[0]) + task33Bytes(s[1:], sym)
}

func main() {
	var str string
	var symStr string
	fmt.Scan(&str)
	fmt.Scan(&symStr)

	if len(symStr) == 0 {
		fmt.Println(str)
		return
	}
	fmt.Println(task33Bytes(str, symStr[0]))
}
