package commands

import (
	"strconv"
	"strings"
	"unicode"
)

func ConvertLastWordsToDecimal(words []string, base int, count int) []string {
	if len(words) == 0 {
		return words
	}

	// Determine how many words to convert (don't exceed available words)
	wordsToConvert := min(count, len(words))

	// Convert the last 'wordsToConvert' words
	startIndex := len(words) - wordsToConvert
	for i := startIndex; i < len(words); i++ {
		originalWord := words[i]

		// Extract the numeric part and surrounding structure
		numericPart, prefix, suffix := extractNumericPart(originalWord)

		// Handle hex prefixes (0x, 0X) for hex conversion
		if base == 16 {
			if strings.HasPrefix(strings.ToLower(numericPart), "0x") {
				numericPart = numericPart[2:] // Remove "0x" or "0X" prefix
			}
		}

		// Handle binary prefixes (0b, 0B) for binary conversion
		if base == 2 {
			if strings.HasPrefix(strings.ToLower(numericPart), "0b") {
				numericPart = numericPart[2:] // Remove "0b" or "0B" prefix
			}
		}

		n, err := strconv.ParseInt(numericPart, base, 64)
		if err != nil {
			// If conversion fails, leave the word unchanged
			continue
		}

		// Reconstruct the word with the converted number
		words[i] = prefix + strconv.FormatInt(n, 10) + suffix
	}

	return words
}

// extractNumericPart extracts the numeric part from a word that might be wrapped in parentheses
// and returns the numeric part, prefix (opening parens), and suffix (closing parens + punctuation)
func extractNumericPart(word string) (numeric, prefix, suffix string) {
	runes := []rune(word)
	if len(runes) == 0 {
		return "", "", ""
	}

	// Find the start of the numeric part (skip opening parentheses)
	start := 0
	for start < len(runes) && runes[start] == '(' {
		start++
	}

	// Find the end of the numeric part (find last alphanumeric character)
	end := len(runes) - 1
	for end >= start && !unicode.IsLetter(runes[end]) && !unicode.IsNumber(runes[end]) {
		end--
	}

	if end < start {
		// No alphanumeric characters found
		return "", "", word
	}

	// Extract parts
	prefix = string(runes[:start])
	numeric = string(runes[start : end+1])
	suffix = string(runes[end+1:])

	return numeric, prefix, suffix
}
