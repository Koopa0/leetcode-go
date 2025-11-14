package CountandSay

import "testing"

func TestCountAndSay(t *testing.T) {
	// Define test cases using a table-driven testing approach
	tests := []struct {
		name     string
		input    int
		expected string
		comment  string
	}{
		{
			name:     "Base case n=1",
			input:    1,
			expected: "1",
			comment:  "First term is defined as '1'",
		},
		{
			name:     "Second term n=2",
			input:    2,
			expected: "11",
			comment:  "Second term describes the first term: one 1",
		},
		{
			name:     "Third term n=3",
			input:    3,
			expected: "21",
			comment:  "Third term describes the second term: two 1s",
		},
		{
			name:     "Fourth term n=4",
			input:    4,
			expected: "1211",
			comment:  "Fourth term describes the third term: one 2, one 1",
		},
		{
			name:     "Fifth term n=5",
			input:    5,
			expected: "111221",
			comment:  "Fifth term describes the fourth term: one 1, one 2, two 1s",
		},
		{
			name:     "Sixth term n=6",
			input:    6,
			expected: "312211",
			comment:  "Sixth term describes the fifth term: three 1s, two 2s, one 1",
		},
		{
			name:     "Edge case - minimum input",
			input:    1,
			expected: "1",
			comment:  "Verifying the minimum allowed input n=1",
		},
		{
			name:     "Medium case n=10",
			input:    10,
			expected: "13211311123113112211",
			comment:  "Testing with a medium sized input",
		},
		// The maximum allowed input n=30 would produce a very long string,
		// so we're not testing the exact output but ensuring it runs without errors
		{
			name:    "Edge case - maximum input",
			input:   30,
			comment: "Test with the maximum allowed input to verify no runtime errors",
		},
	}

	// Execute each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := countAndSay(tc.input)

			// For the maximum input case, we just verify it runs without panicking
			if tc.expected != "" && result != tc.expected {
				t.Errorf("Failed %s: expected %s, got %s", tc.name, tc.expected, result)
			}
		})
	}
}

// Testing performance for larger inputs
func BenchmarkCountAndSay(b *testing.B) {
	// Run the benchmark with n=20
	for i := 0; i < b.N; i++ {
		countAndSay(20)
	}
}
