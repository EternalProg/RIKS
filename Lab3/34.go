package main

import "fmt"

/*
34.Користувач вводить два числа. Напишіть рекурсивну функцію, яка
повертає їх найменше загальне кратне. Отриманий результат виведіть в
термінал.
*/

func gcd(a int, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}

	return gcd(b, a%b)
}

func task34(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	g := gcd(a, b)
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	return a / g * b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(task34(a, b))
}
