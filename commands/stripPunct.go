package commands

import "unicode"

func StripPunctuation(word string) (base, punct string) {
	runes := []rune(word)
	if len(runes) == 0 {
		return "", ""
	}

	// First check if the entire word is wrapped in parentheses
	if len(runes) >= 2 && runes[0] == '(' && runes[len(runes)-1] == ')' {
		// Extract the content inside parentheses
		inside := string(runes[1 : len(runes)-1])
		// Recursively strip punctuation from inside content
		innerBase, innerPunct := StripPunctuation(inside)
		return innerBase, "(" + innerPunct + ")"
	}

	// Handle trailing punctuation
	lastAlphaNum := -1
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsLetter(runes[i]) || unicode.IsNumber(runes[i]) {
			lastAlphaNum = i
			break
		}
	}

	if lastAlphaNum == -1 {
		return "", word
	}

	// Split into main part and trailing punctuation
	mainPart := string(runes[:lastAlphaNum+1])
	trailingPunct := string(runes[lastAlphaNum+1:])

	return mainPart, trailingPunct
}
