package main

import "fmt"

// 50. Продемонструйте використання defer у циклі: чому це може бути проблемою і як це виправити.

func badDeferInLoop() {
	for i := 0; i < 3; i++ {
		defer fmt.Println("cleanup", i)
	}
	fmt.Println("end of loop (defers not run yet)")
}

func fixedDeferInLoop() {
	for i := 0; i < 3; i++ {
		func(i int) {
			defer fmt.Println("cleanup", i)
			fmt.Println("work", i)
		}(i)
	}
}
