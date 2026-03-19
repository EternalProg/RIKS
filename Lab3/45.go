package main

import "fmt"

/*
45.На вхід функції поступає один із символів «>», «<», «=». Використовуючи
механізм замикань порівняйте два значення, які подаються на вхід функції,
що повертається. В результаті повинно повертатися true, або false. У
випадку, коли на вхід основної функції подається невідомий символ –
результат повинен бути false. Отриманий результат виведіть в термінал.
*/

func task45(op string) func(int, int) bool {
	return func(a int, b int) bool {
		switch op {
		case ">":
			return a > b
		case "<":
			return a < b
		case "=":
			return a == b
		default:
			return false
		}
	}
}

func main() {
	var op string
	var a, b int
	fmt.Scan(&op, &a, &b)

	compare := task45(op)
	fmt.Println(compare(a, b))
}
