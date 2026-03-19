package main

import "fmt"

/*
38.Користувач вводить зріз цілих чисел. Напишіть рекурсивну функцію, яка
повертає суму від’ємних елементів зрізу. Отриманий результат виведіть в
термінал.
*/

func task38(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if nums[0] < 0 {
		return nums[0] + task38(nums[1:])
	}
	return task38(nums[1:])
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(task38(nums))
}
