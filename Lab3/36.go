package main

import "fmt"

/*
36.Користувач вводить рядок, який містить одну пару з відкритої та закритої
дужки. Напишіть рекурсивну функцію, яка повертає рядок, який
складається із символів, які знаходяться в дужках вхідного рядка.
Отриманий результат виведіть в термінал.
*/

func task36Rec(s string, i int, started bool) string {
	if i >= len(s) {
		return ""
	}

	ch := s[i]
	if ch == '(' {
		return task36Rec(s, i+1, true)
	}
	if ch == ')' {
		return ""
	}
	if started {
		return string(ch) + task36Rec(s, i+1, true)
	}

	return task36Rec(s, i+1, false)
}

func task36(s string) string {
	return task36Rec(s, 0, false)
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(task36(s))
}
