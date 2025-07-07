package processes

import (
	"strings"
)

func ProcessTextWithGlobalCommand(text string, globalCommand string) string {
	text = normalizeSpaces(text)
	text = CleanSpacesAndPunctuation(text)

	supportedCommands := map[string]bool{
		"up": true, "low": true, "cap": true,
		"hex": true, "bin": true, "rev": true,
		"len": true, "pal": true,
	}

	if !supportedCommands[globalCommand] {
		return text
	}

	// Split text into segments by newlines to preserve structure
	segments := strings.Split(text, "\n")

	for i := 0; i < len(segments); i++ {
		segment := segments[i]
		words := strings.Fields(segment)

		if len(words) > 0 {
			words = ApplyCommand(words, globalCommand, len(words))
		}

		segments[i] = strings.Join(words, " ")
	}

	return strings.Join(segments, "\n")
}
