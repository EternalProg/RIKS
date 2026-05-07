package main

import (
	"fmt"
	"os"
	"strings"
)

// Напишіть програму, яка отримує вміст декількох файлів, об’єднує їх і записує результат в новий файл.

func task1_2() error {
	inputFiles := []string{
		"data/task1/file1.txt",
		"data/task1/file2.txt",
		"data/task1/file3.txt",
	}
	outputFile := "output/task1_2_merged.txt"

	var builder strings.Builder
	for i, path := range inputFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("не вдалося прочитати %s: %w", path, err)
		}

		builder.Write(content)
		if i < len(inputFiles)-1 {
			builder.WriteString("\n")
		}
	}

	if err := os.WriteFile(outputFile, []byte(builder.String()), 0o644); err != nil {
		return fmt.Errorf("не вдалося записати %s: %w", outputFile, err)
	}

	fmt.Println("[1.2] Файли об'єднано в:", outputFile)
	return nil
}
