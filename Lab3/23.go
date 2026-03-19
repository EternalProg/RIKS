package main

import "fmt"

/*
23.Користувач вводить рядок, що містить довільну кількість відкритих та
закритих дужок. Напишіть функцію, яка повертає true або false залежно від
того, чи збережений баланс відкритих і закритих дужок. Якщо у рядку
дужки відсутні, функція повинна повертати true. Отриманий результат
виведіть у термінал.
*/

func task23(s string) bool {
	balance := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			balance++
		} else if s[i] == ')' {
			balance--
			if balance < 0 {
				return false
			}
		}
	}

	return balance == 0
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(task23(s))
}
