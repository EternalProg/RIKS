package main

import "fmt"

// Користувач вводить номер місяця, а код повинен вивести в терміналкількість днів в ньому.

func main() {
	var month int
	if _, err := fmt.Scan(&month); err != nil {
		return
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 4, 6, 9, 11:
		fmt.Println(30)
	case 2:
		fmt.Println(28)
	default:
		fmt.Println("Error")
	}
}
