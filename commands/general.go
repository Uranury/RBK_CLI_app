package commands

import "unicode"

func transformAlnumInPlace(word string, transform func(string) string) string {
	runes := []rune(word)
	var alnums []rune

	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			alnums = append(alnums, r)
		}
	}

	transformed := []rune(transform(string(alnums)))

	index := 0
	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			runes[i] = transformed[index]
			index++
		}
	}

	return string(runes)
}

func applyToWords(words []string, count int, transform func(string) string) []string {
	wordsToConvert := min(count, len(words))

	for i := len(words) - wordsToConvert; i < len(words); i++ {
		words[i] = transform(words[i])
	}

	return words
}
