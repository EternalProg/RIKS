package main

import "fmt"

/*
14.Користувач вводить число n. Напишіть функцію, яка повертає суму
значень від нуля до n-1. Отриманий результат виведіть в термінал.
*/

func task14(n int) int {
	if n == 0 {
		return 0
	}

	return n + task14(n-1)
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(task14(n - 1))
}
