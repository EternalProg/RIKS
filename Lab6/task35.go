package main

import "fmt"

// 35. Реалізуйте стек викликів, де panic виникає на 3-му рівні, а recover — на 1-му. Виведіть порядок виконання defer.

func stackLevel1() {
	defer fmt.Println("defer level1")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	stackLevel2()
}

func stackLevel2() {
	defer fmt.Println("defer level2")
	stackLevel3()
}

func stackLevel3() {
	defer fmt.Println("defer level3")
	panic("panic at level3")
}
