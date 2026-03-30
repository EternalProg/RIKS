package main

import "testing"

func TestStockValueBasic(t *testing.T) {
	stock := Stock{Shares: 10, Quote: 5.5, CurrencyRate: 40}
	expected := 10 * 5.5 * 40
	if stock.Value() != expected {
		t.Fatalf("expected %.2f got %.2f", expected, stock.Value())
	}
}

func TestStockValueInvalidData(t *testing.T) {
	cases := []Stock{
		{Shares: -1, Quote: 10, CurrencyRate: 40},
		{Shares: 10, Quote: -1, CurrencyRate: 40},
		{Shares: 10, Quote: 10, CurrencyRate: 0},
	}
	for _, c := range cases {
		if c.Value() != 0 {
			t.Fatalf("expected 0 for %+v", c)
		}
	}
}

func TestPortfolioTotalValue(t *testing.T) {
	portfolio := InvestmentPortfolio{}
	portfolio.AddStock(Stock{Shares: 10, Quote: 10, CurrencyRate: 40})
	portfolio.AddStock(Stock{Shares: 5, Quote: 20, CurrencyRate: 40})
	expected := float64(10*10*40 + 5*20*40)
	if portfolio.TotalValue() != expected {
		t.Fatalf("expected %.2f got %.2f", expected, portfolio.TotalValue())
	}
}

func TestPortfolioTotalValueByIndustry(t *testing.T) {
	portfolio := InvestmentPortfolio{}
	portfolio.AddStock(Stock{Industry: "Tech", Shares: 10, Quote: 10, CurrencyRate: 1})
	portfolio.AddStock(Stock{Industry: "Agriculture", Shares: 5, Quote: 20, CurrencyRate: 1})
	portfolio.AddStock(Stock{Industry: "TECH", Shares: 2, Quote: 100, CurrencyRate: 1})
	expected := float64(10*10 + 2*100)
	if portfolio.TotalValueByIndustry("tech") != expected {
		t.Fatalf("expected %.2f got %.2f", expected, portfolio.TotalValueByIndustry("tech"))
	}
}

func TestPortfolioStocksByIndustry(t *testing.T) {
	portfolio := InvestmentPortfolio{}
	portfolio.AddStock(Stock{Industry: "Energy", Shares: 10, Quote: 10, CurrencyRate: 1})
	portfolio.AddStock(Stock{Industry: "Energy", Shares: 2, Quote: 5, CurrencyRate: 1})
	portfolio.AddStock(Stock{Industry: "Healthcare", Shares: 3, Quote: 7, CurrencyRate: 1})

	result := portfolio.StocksByIndustry("energy")
	if len(result) != 2 {
		t.Fatalf("expected 2 stocks, got %d", len(result))
	}
}

func TestPortfolioStringContainsTotal(t *testing.T) {
	portfolio := InvestmentPortfolio{}
	portfolio.AddStock(Stock{Company: "A", Shares: 1, Quote: 10, CurrencyRate: 1})
	text := portfolio.String()
	if text == "" || text == "InvestmentPortfolio:" {
		t.Fatalf("expected non-empty string output")
	}
}
