package main

import (
	"fmt"
	"strconv"
)

// Користувач вводить значення. Визначте чи є воно числом, або ні. Виведіть в термінал «Number», або «Other» в залежності від результату перевірки.

func main() {
	var value string
	if _, err := fmt.Scan(&value); err != nil {
		return
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		fmt.Println("Number")
	} else {
		fmt.Println("Other")
	}
}
