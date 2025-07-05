package commands

import (
	"strings"
	"unicode"
)

func capitalizeWord(s string) string {
	runes := []rune(s)
	foundFirstAlpha := false

	for i, r := range runes {
		if unicode.IsLetter(r) {
			if !foundFirstAlpha {
				// Capitalize first letter
				runes[i] = unicode.ToUpper(r)
				foundFirstAlpha = true
			} else {
				runes[i] = unicode.ToLower(r)
			}
		}
	}
	return string(runes)
}

func ApplyTextTransformation(words []string, action string, count int) []string {
	if len(words) == 0 {
		return words
	}

	// Don't transform more words than we have
	wordsToTransform := min(count, len(words))
	startIndex := len(words) - wordsToTransform

	for i := startIndex; i < len(words); i++ {
		base, punct := StripPunctuation(words[i])

		switch action {
		case "cap":
			base = capitalizeWord(base)
		case "up":
			base = strings.ToUpper(base)
		case "low":
			base = strings.ToLower(base)
		}

		words[i] = base + punct
	}

	return words
}
