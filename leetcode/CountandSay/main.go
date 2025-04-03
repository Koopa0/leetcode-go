package CountandSay

import (
	"strconv"
	"strings"
)

// countAndSay returns the nth term of the count-and-say sequence.
// For n = 1, the sequence starts with "1".
// For n > 1, each term is derived by "counting and saying" the previous term.
func countAndSay(n int) string {
	// Base case: the first term is "1"
	if n == 1 {
		return "1"
	}

	// Start with the first term
	prev := "1"

	// Generate terms iteratively from 2 to n
	for i := 2; i <= n; i++ {
		// Use strings.Builder for efficient string concatenation
		var result strings.Builder

		// Initialize counter and character from the first character
		count := 1
		char := prev[0]

		// Process the previous term character by character
		for j := 1; j < len(prev); j++ {
			if prev[j] == char {
				// If the current character is the same as previous, increment count
				count++
			} else {
				// If the current character is different, append count+char to result
				result.WriteString(strconv.Itoa(count))
				result.WriteByte(char)

				// Reset character and count for the new group
				char = prev[j]
				count = 1
			}
		}

		// Don't forget to append the count and character for the last group
		result.WriteString(strconv.Itoa(count))
		result.WriteByte(char)

		// Update the previous term to the newly generated term
		prev = result.String()
	}

	return prev
}
