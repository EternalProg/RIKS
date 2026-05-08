package main

import "fmt"

// Напишіть програму для знаходження з використанням горутин суми середніх арифметичних значень елементів трьох зрізів.

func task7SumOfMeans() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7}
	c := []int{8, 9}

	ch := make(chan float64, 3)

	avg := func(s []int) {
		if len(s) == 0 {
			ch <- 0
			return
		}
		sum := 0
		for _, v := range s {
			sum += v
		}
		ch <- float64(sum) / float64(len(s))
	}

	go avg(a)
	go avg(b)
	go avg(c)

	total := (<-ch) + (<-ch) + (<-ch)
	fmt.Printf("Сума середніх арифметичних = %.4f\n", total)
}
