package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Вам надано файл, який містить дані в JSON-форматі. Напишіть програму, яка дозволяє здійснити з нього завантаження, змінити деякі значення і зберегти поточний стан екземплярів класів у новий файл в JSON-форматі.

type Property struct {
	Address string `json:"address"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Rental struct {
	Type        string   `json:"type"`
	Bedrooms    int      `json:"bedrooms"`
	Bathrooms   float64  `json:"bathrooms"`
	Amenities   []string `json:"amenities"`
	PetsAllowed bool     `json:"pets_allowed"`
}

type Landlord struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Review struct {
	Username string  `json:"username"`
	Rating   int     `json:"rating"`
	Comment  *string `json:"comment,omitempty"`
	Date     string  `json:"date"`
}

type Lease struct {
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	MonthlyRent float64 `json:"monthly_rent"`
}

type RentalData struct {
	Property Property `json:"property"`
	Rental   Rental   `json:"rental"`
	Landlord Landlord `json:"landlord"`
	Reviews  []Review `json:"reviews"`
	Lease    *Lease   `json:"lease"`
}

func task2JSON() error {
	inputFile := "data/task2/rental_input.json"
	outputFile := "output/task2_updated.json"

	raw, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("не вдалося прочитати %s: %w", inputFile, err)
	}

	var data RentalData
	if err := json.Unmarshal(raw, &data); err != nil {
		return fmt.Errorf("помилка розбору JSON: %w", err)
	}

	data.Property.City = "Liverpool"
	data.Rental.Bedrooms = 3
	data.Rental.PetsAllowed = false
	data.Rental.Amenities = append(data.Rental.Amenities, "Wi-Fi")
	data.Landlord.Phone = "+1 987-654-3210"
	data.Lease = &Lease{
		StartDate:   "2024-09-01",
		EndDate:     "2025-08-31",
		MonthlyRent: 1800,
	}

	updated, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("помилка серіалізації JSON: %w", err)
	}

	if err := os.WriteFile(outputFile, updated, 0o644); err != nil {
		return fmt.Errorf("не вдалося записати %s: %w", outputFile, err)
	}

	fmt.Println("\n[2] JSON завантажено, змінено і збережено у:", outputFile)
	return nil
}
