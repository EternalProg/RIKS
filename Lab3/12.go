package main

import "fmt"

/*
12.Користувач вводить 2 числа: A та B. Напишіть функцію, яка буде
перевіряти чи у A встановлений біт під номером B в одиницю, чи ні.
Результат перевірки необхідно повернути у вигляді: true – так, ні – false, і
вивести його в термінал.
*/

func task12(A int, B int) bool {
	return A&(1<<B) != 0
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	fmt.Println(task12(A, B))
}
