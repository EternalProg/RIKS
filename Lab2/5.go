package main

import "fmt"

/*
Користувач вводить свій оклад до вирахування податків.
Якщо він більше за 50 000 гривень,
то вирахуйте з нього 13% ПДФО, в іншому випадку – 6%.
Виведіть отриманий результат в термінал. Реалізуйте 2 версії коду з різними підходами.
*/

func main() {
	var salary float64
	if _, err := fmt.Scan(&salary); err != nil {
		return
	}

	var resIf float64
	if salary > 50000 {
		resIf = salary * 0.87
	} else {
		resIf = salary * 0.94
	}

	var resSwitch float64
	switch {
	case salary > 50000:
		resSwitch = salary * 0.87
	default:
		resSwitch = salary * 0.94
	}

	fmt.Println(resIf)
	fmt.Println(resSwitch)
}
