package main

import "errors"

// 5. Реалізуйте ланцюжок з трьох функцій, кожна з яких може повернути помилку. Виконайте послідовну перевірку помилок у main.

func step1(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("step1: n must be positive")
	}
	return n, nil
}

func step2(n int) (int, error) {
	if n%2 != 0 {
		return 0, errors.New("step2: n must be even")
	}
	return n, nil
}

func step3(n int) (int, error) {
	if n >= 100 {
		return 0, errors.New("step3: n must be less than 100")
	}
	return n, nil
}
