package processes

import (
	"CLI_app/commands"
	"regexp"
	"strconv"
	"strings"
)

func ApplyCommand(words []string, action string, count int) []string {
	if len(words) == 0 {
		return words
	}

	switch action {
	case "up", "low", "cap":
		return commands.ApplyTextTransformation(words, action, count)
	case "hex":
		return commands.ConvertLastWordsToDecimal(words, 16, count)
	case "bin":
		return commands.ConvertLastWordsToDecimal(words, 2, count)
	case "rev":
		return commands.ReverseTheWord(words, count)
	case "len":
		return commands.ReplaceTheWordToItsLen(words, count)
	case "pal":
		return commands.ReplaceWordToPal(words, count)
	default:
		return words
	}
}

func CleanSpacesAndPunctuation(text string) string {
	// Remove space before standard punctuation
	reBeforePunct := regexp.MustCompile(`\s+([,.;:?!\)\]}])`)
	text = reBeforePunct.ReplaceAllString(text, "$1")

	// Remove space after opening punctuation
	reAfterOpen := regexp.MustCompile(`([(\[{])\s+`)
	text = reAfterOpen.ReplaceAllString(text, "$1")

	// Preserve space before opening quotes
	reBeforeQuote := regexp.MustCompile(`(\w)\s+'\s*(\w)`)
	text = reBeforeQuote.ReplaceAllString(text, "$1 '$2")

	// Fix contractions while preserving quoted words
	reContractions := regexp.MustCompile(`(\b\w+)\s+'([msd]|ll|ve|re|t)([^a-zA-Z]|$)`)
	text = reContractions.ReplaceAllString(text, "$1'$2$3")

	return text
}

func ParseCount(countStr string) int {
	if countStr == "" {
		return 1
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 1 // Invalid string like "abc" defaults to 1
	}

	if count <= 0 {
		return 1 // Zero or negative counts default to 1
	}

	if count > 1000 {
		return 1000
	}

	return count
}

func normalizeSpaces(text string) string {
	// Only normalize spaces within lines, keep the newlines
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = regexp.MustCompile(`[^\S\n]+`).ReplaceAllString(lines[i], " ")
		lines[i] = strings.TrimSpace(lines[i])
	}
	return strings.Join(lines, "\n")
}

func reconstructWithSpacing(original string, processedWords []string) string {
	if len(processedWords) == 0 {
		return original
	}

	// Create a regex to find word boundaries
	wordRegex := regexp.MustCompile(`\S+`)
	wordPositions := wordRegex.FindAllStringIndex(original, -1)

	if len(wordPositions) == 0 {
		return original
	}

	result := original
	wordIndex := len(processedWords) - 1

	// Replace words from right to left to maintain correct positions
	for i := len(wordPositions) - 1; i >= 0 && wordIndex >= 0; i-- {
		start := wordPositions[i][0]
		end := wordPositions[i][1]
		result = result[:start] + processedWords[wordIndex] + result[end:]
		wordIndex--
	}

	return result
}

func ProcessText(text string) string {
	commandRegex := regexp.MustCompile(`\((up|low|cap|hex|bin|rev|len|pal)(?:\s*,\s*([^)]+))?\)`)

	for {
		matches := commandRegex.FindStringSubmatch(text)
		if matches == nil {
			break
		}

		action := matches[1]
		count := ParseCount(matches[2])
		commandPos := commandRegex.FindStringIndex(text)
		if commandPos == nil {
			break
		}

		beforeCommand := text[:commandPos[0]]
		afterCommand := text[commandPos[1]:]

		words := strings.Fields(beforeCommand)
		words = ApplyCommand(words, action, count)

		processedBefore := reconstructWithSpacing(beforeCommand, words)
		text = processedBefore + afterCommand
	}

	text = normalizeSpaces(text)
	return CleanSpacesAndPunctuation(text)
}
