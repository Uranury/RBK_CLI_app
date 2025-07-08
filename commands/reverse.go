package commands

func ReverseTheWord(words []string, count int) []string {
	return applyToWords(words, count, func(word string) string {
		return transformAlnumInPlace(word, reverseString)
	})
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
