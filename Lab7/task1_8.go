package main

import (
	"fmt"
	"os"
	"strings"
)

// Напишіть програму, яка отримує вміст текстового файлу і підраховує кількість голосних і приголосних букв. Виведіть в термінал отриманий результат і запишіть в новий файл.

func task1_8() error {
	inputFile := "data/task1/text_for_count.txt"
	outputFile := "output/task1_8_counts.txt"

	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("не вдалося прочитати %s: %w", inputFile, err)
	}

	vowelCount, consonantCount := countVowelsAndConsonants(string(content))
	result := fmt.Sprintf("Кількість голосних: %d\nКількість приголосних: %d\n", vowelCount, consonantCount)

	if err := os.WriteFile(outputFile, []byte(result), 0o644); err != nil {
		return fmt.Errorf("не вдалося записати %s: %w", outputFile, err)
	}

	fmt.Println("\n[1.8]", strings.TrimSpace(result))
	fmt.Println("Результат записано у:", outputFile)

	return nil
}
