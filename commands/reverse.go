package commands

import "unicode"

func ReverseTheWord(words []string, count int) []string {
	wordsToConvert := count
	if wordsToConvert > len(words) {
		wordsToConvert = len(words)
	}

	for i := len(words) - wordsToConvert; i < len(words); i++ {
		words[i] = reverseAlnumInPlace(words[i])
	}

	return words
}

func reverseAlnumInPlace(word string) string {
	runes := []rune(word)
	var letters []rune

	// Extract alphanumeric runes
	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			letters = append(letters, r)
		}
	}

	// Reverse the letters
	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	// Replace the alphanumeric runes in the original rune slice
	index := 0
	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			runes[i] = letters[index]
			index++
		}
	}

	return string(runes)
}
