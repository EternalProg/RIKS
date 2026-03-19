package main

import "fmt"

/*

Користувач вводить 2 значення (x і y). Визначте в якій четверті знаходиться точка з отриманими координатами і виведіть її в термінал (1, 2, 3, 4, або їх пересікання).

*/

func main() {
	var x, y float64
	if _, err := fmt.Scan(&x, &y); err != nil {
		return
	}

	if x == 0 || y == 0 {
		fmt.Println("intersection")
		return
	}

	if x > 0 && y > 0 {
		fmt.Println(1)
	} else if x < 0 && y > 0 {
		fmt.Println(2)
	} else if x < 0 && y < 0 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
