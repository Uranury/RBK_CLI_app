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

	wordsToConvert := min(count, len(words))
	startIndex := len(words) - wordsToConvert

	for i := startIndex; i < len(words); i++ {
		switch action {
		case "cap":
			words[i] = transformAlnumInPlace(words[i], capitalizeWord)
		case "up":
			words[i] = transformAlnumInPlace(words[i], strings.ToUpper)
		case "low":
			words[i] = transformAlnumInPlace(words[i], strings.ToLower)
		}
	}
	return words
}
