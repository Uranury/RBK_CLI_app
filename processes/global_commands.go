package processes

import (
	"strings"
)

func ProcessTextWithGlobalCommand(text string, globalCommand string) string {
	// Validate that it's a supported command
	supportedCommands := map[string]bool{
		"up": true, "low": true, "cap": true,
		"hex": true, "bin": true, "rev": true,
		"len": true, "pal": true,
	}

	if !supportedCommands[globalCommand] {
		// If invalid command, just clean spaces and return
		return CleanSpacesAndPunctuation(normalizeSpaces(text))
	}

	// Split text into segments by newlines to preserve line structure
	segments := strings.Split(text, "\n")

	for i := 0; i < len(segments); i++ {
		segment := segments[i]

		// Get all words from the segment
		words := strings.Fields(segment)

		// Apply the global command to all words in the segment
		if len(words) > 0 {
			words = ApplyCommand(words, globalCommand, len(words))
		}

		// Reconstruct the segment
		segments[i] = strings.Join(words, " ")
	}

	// Join segments back with newlines
	result := strings.Join(segments, "\n")

	// Apply space normalization and punctuation cleanup
	return CleanSpacesAndPunctuation(result)
}
