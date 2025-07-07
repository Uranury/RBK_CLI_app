package commands

import "unicode"

func StripPunctuation(word string) (base, punct string) {
	runes := []rune(word)
	if len(runes) == 0 {
		return "", ""
	}

	// If word is fully wrapped in parentheses, try to handle inside
	if len(runes) >= 2 && runes[0] == '(' && runes[len(runes)-1] == ')' {
		inside := string(runes[1 : len(runes)-1])
		innerBase, innerPunct := StripPunctuation(inside)

		// Only wrap again if innerPunct changed the content
		if innerBase != inside || innerPunct != "" {
			return innerBase, "(" + innerPunct + ")"
		} else {
			// Don't strip punctuation that doesn't need transformation
			return word, ""
		}
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

	mainPart := string(runes[:lastAlphaNum+1])
	trailingPunct := string(runes[lastAlphaNum+1:])

	return mainPart, trailingPunct
}
