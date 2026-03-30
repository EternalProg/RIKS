package main

import (
	"fmt"
)

func main() {
	portfolio := InvestmentPortfolio{}
	portfolio.AddStock(Stock{
		Company:      "TechNova",
		Issuer:       "TechNova Inc.",
		Industry:     "Technology",
		Shares:       120,
		Quote:        15.75,
		CurrencyRate: 41.20,
	})
	portfolio.AddStock(Stock{
		Company:      "AgroLife",
		Issuer:       "AgroLife Ltd",
		Industry:     "Agriculture",
		Shares:       80,
		Quote:        10.10,
		CurrencyRate: 41.20,
	})
	fmt.Println(portfolio.String())

	store := JewelryStore{}
	store.AddItem(Jewelry{Name: "Aurora Ring", Kind: "Ring", Price: 2500})
	store.AddItem(Jewelry{Name: "Ocean Necklace", Kind: "Necklace", Price: 4100})
	store.AddItem(Jewelry{Name: "Twilight Earrings", Kind: "Earrings", Price: 1950})
	store.Sell("Aurora Ring")
	fmt.Println(store.String())
}
