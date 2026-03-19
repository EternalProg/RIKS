package main

import "fmt"

/*

Користувач вводить зріз цілих чисел, який складається мінімум із 9 елементов. Якщо значення в середині зрізу більше, або дорівнює 10, то виведіть в термінал суму його першого та останнього елементів, в іншому випадку добуток першого та останнього.


*/

func main() {
	nums := make([]int, 0)
	for {
		var v int
		if _, err := fmt.Scan(&v); err != nil {
			break
		}
		nums = append(nums, v)
	}

	if len(nums) == 0 || len(nums) < 9 {
		fmt.Println("At least 9 numbers are required")
		return
	}

	mid := len(nums) / 2
	fmt.Println("Middle element:", nums[mid])
	if nums[mid] >= 10 {
		fmt.Println(nums[0] + nums[len(nums)-1])
	} else {
		fmt.Println(nums[0] * nums[len(nums)-1])
	}
}
