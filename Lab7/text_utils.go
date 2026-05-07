package main

import (
	"strings"
	"unicode"
)

var vowels = map[rune]struct{}{
	'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'y': {},
	'а': {}, 'е': {}, 'и': {}, 'і': {}, 'ї': {}, 'о': {}, 'у': {}, 'ю': {}, 'я': {}, 'є': {},
}

func countVowelsAndConsonants(text string) (int, int) {
	vowelsCount := 0
	consonantsCount := 0

	for _, r := range strings.ToLower(text) {
		if !isSupportedLetter(r) {
			continue
		}
		if isVowel(r) {
			vowelsCount++
		} else {
			consonantsCount++
		}
	}

	return vowelsCount, consonantsCount
}

func mostFrequentVowelAndConsonant(text string) (rune, int, rune, int) {
	vowelFrequency := make(map[rune]int)
	consonantFrequency := make(map[rune]int)

	for _, r := range strings.ToLower(text) {
		if !isSupportedLetter(r) {
			continue
		}

		if isVowel(r) {
			vowelFrequency[r]++
		} else {
			consonantFrequency[r]++
		}
	}

	mostVowel, maxVowel := maxRuneFrequency(vowelFrequency)
	mostConsonant, maxConsonant := maxRuneFrequency(consonantFrequency)

	return mostVowel, maxVowel, mostConsonant, maxConsonant
}

func maxRuneFrequency(freq map[rune]int) (rune, int) {
	var topRune rune
	maxCount := 0
	for r, count := range freq {
		if count > maxCount {
			topRune = r
			maxCount = count
		}
	}
	return topRune, maxCount
}

func isVowel(r rune) bool {
	_, ok := vowels[r]
	return ok
}

func isSupportedLetter(r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}

	return isLatinLetter(r) || isCyrillicLetter(r)
}

func isLatinLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isCyrillicLetter(r rune) bool {
	return (r >= 'а' && r <= 'я') || r == 'і' || r == 'ї' || r == 'є' || r == 'ґ'
}
