package main

import (
	"fmt"
	"strings"
)

/*
10.Реалізуйте клас «Акції», який міститиме інформацію про компанію,
емітент акцій, кількість акцій у портфелі, поточне котирування та курс
заданої валюти, а також клас «Інвестиційний портфель». Портфель
інвестора повинен містити методи отримання повної сумарної вартості
всіх акцій у портфелі; з певної галузі; а також повертати список акцій
компанії заданої галузі. Також забезпечте виведення поточного стану
об'єктів у термінал.

*/

type Stock struct {
	Company      string
	Issuer       string
	Industry     string
	Shares       int
	Quote        float64
	CurrencyRate float64
}

func (s Stock) Value() float64 {
	if s.Shares <= 0 || s.Quote < 0 || s.CurrencyRate <= 0 {
		return 0
	}
	return float64(s.Shares) * s.Quote * s.CurrencyRate
}

func (s Stock) String() string {
	return fmt.Sprintf(
		"Stock{Company:%s, Issuer:%s, Industry:%s, Shares:%d, Quote:%.2f, Rate:%.4f, Value:%.2f}",
		s.Company,
		s.Issuer,
		s.Industry,
		s.Shares,
		s.Quote,
		s.CurrencyRate,
		s.Value(),
	)
}

type InvestmentPortfolio struct {
	stocks []Stock
}

func (p *InvestmentPortfolio) AddStock(stock Stock) {
	p.stocks = append(p.stocks, stock)
}

// Отримання повної сумарної вартості портфеля
func (p InvestmentPortfolio) TotalValue() float64 {
	total := 0.0
	for _, s := range p.stocks {
		total += s.Value()
	}
	return total
}

// Отримання повної сумарної вартості портфеля за галузю
func (p InvestmentPortfolio) TotalValueByIndustry(industry string) float64 {
	total := 0.0
	for _, s := range p.stocks {
		if strings.EqualFold(s.Industry, industry) {
			total += s.Value()
		}
	}
	return total
}

func (p InvestmentPortfolio) StocksByIndustry(industry string) []Stock {
	result := []Stock{}
	for _, s := range p.stocks {
		if strings.EqualFold(s.Industry, industry) {
			result = append(result, s)
		}
	}
	return result
}

func (p InvestmentPortfolio) String() string {
	lines := []string{"InvestmentPortfolio:"}
	for _, s := range p.stocks {
		lines = append(lines, "  "+s.String())
	}
	lines = append(lines, fmt.Sprintf("TotalValue: %.2f", p.TotalValue()))
	return strings.Join(lines, "\n")
}
