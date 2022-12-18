package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Create a map mapping keywords to emoji characters
	emoji := map[string]string{
		"smile":      "😃",
		"bear":       "🐻",
		"hamburger":  "🍔",
		"lightbulb":  "💡",
		"idea":       "💡",
		"comment":    "💬",
		"chat":       "💬",
		"pomo":       "🍅",
		"stop":       "🛑",
		"warning":    "⚠️",
		"rant":       "🤬",
		"tv":         "📺",
		"update":     "📰",
		"tux":        "🐧",
		"facepalm":   "🤦",
		"puke":       "🤢",
		"skull":      "💀",
		"wizard":     "🧙",
	}

	// Get the input file name from the command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: emoji_replacer INPUT_FILE")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// Read the input file
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Replace all occurrences of the keywords with the corresponding emoji characters
	output := string(input)
	for keyword, emojiChar := range emoji {
		output = strings.ReplaceAll(output, ":"+keyword+":", emojiChar)
	}

	// Write the modified text to the output file
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

