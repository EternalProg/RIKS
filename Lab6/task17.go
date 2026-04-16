package main

import (
	"errors"
	"fmt"
	"strings"
)

// 17.Реалізуйте тип MultiError, що зберігає зріз помилок. Реалізуйте метод Error() string та функцію validateForm, яка збирає кілька помилок.

type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	if len(m.Errors) == 0 {
		return ""
	}
	var b strings.Builder
	for i, err := range m.Errors {
		if i > 0 {
			b.WriteString("; ")
		}
		b.WriteString(err.Error())
	}
	return b.String()
}

func (m *MultiError) Add(err error) {
	if err != nil {
		m.Errors = append(m.Errors, err)
	}
}

func (m *MultiError) OrNil() error {
	if len(m.Errors) == 0 {
		return nil
	}
	return m
}

func validateForm(name, email string, age int) error {
	m := &MultiError{}
	if strings.TrimSpace(name) == "" {
		m.Add(errors.New("name is required"))
	}
	if !strings.Contains(email, "@") {
		m.Add(errors.New("email must contain @"))
	}
	if age < 0 || age > 100 {
		m.Add(fmt.Errorf("age %d out of range", age))
	}
	return m.OrNil()
}
