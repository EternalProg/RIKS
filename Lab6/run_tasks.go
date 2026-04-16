package main

import (
	"errors"
	"fmt"
)

func runTask5() {
	value, err := step1(15)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	value, err = step2(value)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	value, err = step3(value)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", value)
}

func runTask17() {
	if err := validateForm("", "bad", 200); err != nil {
		fmt.Println("validation error:", err)
	}
}

func runTask25() {
	err := level1()
	fmt.Println("errors.Is ErrLowLevel:", errors.Is(err, ErrLowLevel))
	var we *WrappedError
	if errors.As(err, &we) {
		fmt.Println("errors.As WrappedError:", we.msg)
	}
}

func runTask35() {
	stackLevel1()
}

func runTask50() {
	fmt.Println("bad defer in loop")
	badDeferInLoop()
	fmt.Println("fixed with scope")
	fixedDeferInLoop()
}
