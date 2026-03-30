package main

import (
	"fmt"
	"strings"
)

/*
13.Реалізуйте клас JewelryStore (Ювелірний магазин), який є системою
управління продажі ювелірних виробів. Клас повинен мати методи для
додавання нового виробу, видалення виробу, пошуку виробу за назвою,
списку виробів де ціна >= заданої та розрахунку загальної вартості всіх
виробів у магазині. У системі мають бути передбачені різні типи
ювелірних виробів (наприклад, кільця, намиста, сережки) з різними цінами
та властивостями. Клас JewelryStore повинен також надавати методи для
виведення поточного стану виробів та інформації про продаж терміналу.
*/

type Jewelry struct {
	Name  string
	Kind  string
	Price float64
}

func (j Jewelry) String() string {
	return fmt.Sprintf("Jewelry{Name:%s, Kind:%s, Price:%.2f}", j.Name, j.Kind, j.Price)
}

type JewelryStore struct {
	items []Jewelry
	sales float64
}

func (s *JewelryStore) AddItem(item Jewelry) {
	if item.Price < 0 {
		return
	}
	s.items = append(s.items, item)
}

func (s *JewelryStore) RemoveItem(name string) bool {
	for i, item := range s.items {
		if strings.EqualFold(item.Name, name) {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return true
		}
	}
	return false
}

func (s JewelryStore) FindByName(name string) (Jewelry, bool) {
	for _, item := range s.items {
		if strings.EqualFold(item.Name, name) {
			return item, true
		}
	}
	return Jewelry{}, false
}

func (s JewelryStore) ItemsByMinPrice(minPrice float64) []Jewelry {
	result := []Jewelry{}
	for _, item := range s.items {
		if item.Price >= minPrice {
			result = append(result, item)
		}
	}
	return result
}

func (s JewelryStore) TotalValue() float64 {
	total := 0.0
	for _, item := range s.items {
		total += item.Price
	}
	return total
}

func (s *JewelryStore) Sell(name string) bool {
	item, ok := s.FindByName(name)
	if !ok {
		return false
	}
	s.sales += item.Price
	s.RemoveItem(name)
	return true
}

func (s JewelryStore) Sales() float64 {
	return s.sales
}

func (s JewelryStore) String() string {
	lines := []string{"JewelryStore:"}
	for _, item := range s.items {
		lines = append(lines, "  "+item.String())
	}
	lines = append(lines, fmt.Sprintf("TotalValue: %.2f", s.TotalValue()))
	lines = append(lines, fmt.Sprintf("Sales: %.2f", s.sales))
	return strings.Join(lines, "\n")
}
