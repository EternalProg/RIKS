package main

import (
	"fmt"
	"os"
	"strings"
)

// Напишіть програму, яка отримує вміст текстового файлу, записує в новий файл і виводить в термінал голосну і приголосну букви, які найчастіше зустрічаються.

func task1_12() error {
	inputFile := "data/task1/text_for_top_letters.txt"
	outputFile := "output/task1_12_top_letters.txt"

	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("не вдалося прочитати %s: %w", inputFile, err)
	}

	mostVowel, vowelFreq, mostConsonant, consonantFreq := mostFrequentVowelAndConsonant(string(content))

	result := fmt.Sprintf(
		"Найчастіша голосна: %q (кількість: %d)\nНайчастіша приголосна: %q (кількість: %d)\n",
		string(mostVowel), vowelFreq, string(mostConsonant), consonantFreq,
	)

	if err := os.WriteFile(outputFile, []byte(result), 0o644); err != nil {
		return fmt.Errorf("не вдалося записати %s: %w", outputFile, err)
	}

	fmt.Println("\n[1.12]", strings.TrimSpace(result))
	fmt.Println("Результат записано у:", outputFile)

	return nil
}
