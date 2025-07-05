package commands

import (
	"strconv"
	"unicode"
)

func ReplaceTheWordToItsLen(words []string, count int) []string {
	wordsToConvert := count
	if wordsToConvert > len(words) {
		wordsToConvert = len(words)
	}

	for i := len(words) - wordsToConvert; i < len(words); i++ {
		words[i] = replaceBaseWithTotalCount(words[i])
	}

	return words
}

func replaceBaseWithTotalCount(word string) string {
	runes := []rune(word)
	totalBase := 0
	firstBaseIndex := -1

	// Count base characters and find first base position
	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			totalBase++
			if firstBaseIndex == -1 {
				firstBaseIndex = i
			}
		}
	}

	if totalBase == 0 {
		return word // No base characters found
	}

	countStr := strconv.Itoa(totalBase)
	var newRunes []rune
	inserted := false

	// Build new string with punctuation in original positions
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if !inserted {
				newRunes = append(newRunes, []rune(countStr)...)
				inserted = true
			}
			// Skip other base characters
		} else {
			newRunes = append(newRunes, r)
		}
	}

	return string(newRunes)
}
