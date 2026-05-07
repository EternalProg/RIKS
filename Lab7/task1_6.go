package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Напишіть програму, яка отримує вміст двох текстових файлів і порівнює їх, записуючи в новий файл і виводячи в термінал рядки, які є тільки в одному із файлів.

func task1_6() error {
	fileA := "data/task1/compare1.txt"
	fileB := "data/task1/compare2.txt"
	outputFile := "output/task1_6_diff.txt"

	linesA, err := readLines(fileA)
	if err != nil {
		return err
	}
	linesB, err := readLines(fileB)
	if err != nil {
		return err
	}

	setA := make(map[string]struct{}, len(linesA))
	setB := make(map[string]struct{}, len(linesB))
	for _, line := range linesA {
		setA[line] = struct{}{}
	}
	for _, line := range linesB {
		setB[line] = struct{}{}
	}

	uniqueA := make([]string, 0)
	for _, line := range linesA {
		if _, existsInB := setB[line]; existsInB {
			continue
		}
		uniqueA = append(uniqueA, line)
	}

	uniqueB := make([]string, 0)
	for _, line := range linesB {
		if _, existsInA := setA[line]; existsInA {
			continue
		}
		uniqueB = append(uniqueB, line)
	}

	var out strings.Builder
	out.WriteString("Рядки, що є тільки у першому файлі:\n")
	for _, line := range uniqueA {
		out.WriteString(line + "\n")
	}
	out.WriteString("\nРядки, що є тільки у другому файлі:\n")
	for _, line := range uniqueB {
		out.WriteString(line + "\n")
	}

	if err := os.WriteFile(outputFile, []byte(out.String()), 0o644); err != nil {
		return fmt.Errorf("не вдалося записати %s: %w", outputFile, err)
	}

	fmt.Println("\n[1.6] Рядки, що є тільки в одному з файлів:")
	for _, line := range uniqueA {
		fmt.Printf("Тільки у %s: %s\n", fileA, line)
	}
	for _, line := range uniqueB {
		fmt.Printf("Тільки у %s: %s\n", fileB, line)
	}
	fmt.Println("Результат також записано у:", outputFile)

	return nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("не вдалося відкрити %s: %w", path, err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("помилка читання %s: %w", path, err)
	}

	return lines, nil
}
