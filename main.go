package main

import (
	"CLI_app/processes"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 || (len(args) == 2 && (args[1] == "--help" || args[1] == "-h")) {
		PrintHomePage()
		return
	}

	if len(args) < 3 || len(args) > 4 {
		fmt.Println("Usage: go run . <input.txt> <output.txt> [global_command]")
		fmt.Println("Use --help or -h for detailed information.")
		os.Exit(1)
	}

	inputFile := args[1]
	outputFile := args[2]
	var globalCommand string
	if len(args) == 4 {
		globalCommand = args[3]
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Failed to read %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	fmt.Println("Processing...")

	var processed string
	if globalCommand != "" {
		processed = processes.ProcessTextWithGlobalCommand(string(data), globalCommand)
	} else {
		processed = processes.ProcessText(string(data))
	}

	err = os.WriteFile(outputFile, []byte(processed), 0644)
	if err != nil {
		fmt.Printf("❌ Failed to write %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Println("Done✅")
}
