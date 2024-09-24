package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Specify the path to your text file
	filePath := "myfile.txt" // Replace with your file path

	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert content to string and remove backslashes
	modifiedContent := string(content)
	modifiedContent = removeBackslashes(modifiedContent)

	// Remove quotes based on conditions
	modifiedContent = removeQuotes(modifiedContent)

	// Write the modified content back to the file
	err = ioutil.WriteFile(filePath, []byte(modifiedContent), 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Println("Backslashes and unwanted quotes removed; content rewritten successfully.")
}

// removeBackslashes removes all backslashes from the input string
func removeBackslashes(s string) string {
	return strings.ReplaceAll(s, "\\", "")
}

// removeQuotes removes double quotes based on specified conditions
func removeQuotes(s string) string {
	var result strings.Builder
	length := len(s)

	for i := 0; i < length; i++ {
		// Check for the conditions to remove quotes
		if s[i] == '"' {
			if i > 0 && s[i-1] == '}' { // Remove quote before '{'
				continue
			} else if i < length-1 && s[i+1] == '{' { // Remove quote after '}'
				continue
			}
		}
		// Append the character to the result
		result.WriteByte(s[i])
	}

	return result.String()
}
