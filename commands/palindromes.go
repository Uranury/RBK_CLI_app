package commands

import "unicode"

func ReplaceWordToPal(words []string, count int) []string {
	wordsToConvert := min(count, len(words))

	for i := len(words) - wordsToConvert; i < len(words); i++ {
		words[i] = replacePalindromeInPlace(words[i])
	}

	return words
}

func replacePalindromeInPlace(word string) string {
	runes := []rune(word)

	// Find the first and last alphanumeric index
	start, end := -1, -1
	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if start == -1 {
				start = i
			}
			end = i
		}
	}

	if start == -1 {
		return word // no alphanumeric chars
	}

	core := string(runes[start : end+1])
	if !checkIfWordIsPalindrome(core) {
		return word
	}

	// Reconstruct word with "pal" in place of core
	return string(runes[:start]) + "pali" + string(runes[end+1:])
}

/*
func checkIfWordIsPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
*/

func checkIfWordIsPalindrome(word string) bool {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
