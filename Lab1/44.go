package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Завдання №44
// Користувач вводить з клавіатури 2 значення val1, val2.
// Порахувати їх добуток і вивести отриманий результат в термінал.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter val1: ")
	val1Input, _ := reader.ReadString('\n')
	val1, _ := strconv.ParseFloat(strings.TrimSpace(val1Input), 64)

	fmt.Print("Enter val2: ")
	val2Input, _ := reader.ReadString('\n')
	val2, _ := strconv.ParseFloat(strings.TrimSpace(val2Input), 64)

	product := val1 * val2

	fmt.Printf("Product: %f\n", product)
}
