package main

import (
	"errors"
	"fmt"
)

// 25.Реалізуйте власний тип помилки з методом Unwrap(). Продемонструйте роботу errors.Is та errors.As через ланцюжок загортань.

var ErrLowLevel = errors.New("low level failure")

type WrappedError struct {
	msg string
	err error
}

func (e *WrappedError) Error() string {
	if e.err == nil {
		return e.msg
	}
	return fmt.Sprintf("%s: %v", e.msg, e.err)
}

func (e *WrappedError) Unwrap() error {
	return e.err
}

func level3() error {
	return ErrLowLevel
}

func level2() error {
	if err := level3(); err != nil {
		return &WrappedError{msg: "level2", err: err}
	}
	return nil
}

func level1() error {
	if err := level2(); err != nil {
		return &WrappedError{msg: "level1", err: err}
	}
	return nil
}
