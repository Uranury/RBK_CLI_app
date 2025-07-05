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

	// Preserve space before opening quotesa
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

func processSegment(segment string) string {
	commandRegex := regexp.MustCompile(`\((up|low|cap|hex|bin|rev|len|pal)(?:\s*,\s*([^)]+))?\)`)

	for {
		matches := commandRegex.FindStringSubmatch(segment)
		if matches == nil {
			break
		}

		action := matches[1]
		count := ParseCount(matches[2])
		commandPos := commandRegex.FindStringIndex(segment)
		if commandPos == nil {
			break
		}

		beforeCommand := segment[:commandPos[0]]
		afterCommand := segment[commandPos[1]:]

		// Preserve existing newlines within the segment
		words := strings.Fields(beforeCommand)
		words = ApplyCommand(words, action, count)

		// Reconstruct segment while preserving original spacing structure
		segment = strings.Join(words, " ") + afterCommand
	}

	// Normalize spaces but preserve newlines
	segment = normalizeSpaces(segment)

	// Apply punctuation cleanup
	return CleanSpacesAndPunctuation(segment)
}

func normalizeSpaces(text string) string {
	return regexp.MustCompile(`[^\S\n]+`).ReplaceAllString(text, " ")
}

func ProcessText(text string) string {
	// Preserve newlines by splitting and processing segments
	segments := strings.Split(text, "\n")
	for i := 0; i < len(segments); i++ {
		segments[i] = processSegment(segments[i])
	}
	return strings.Join(segments, "\n")
}
