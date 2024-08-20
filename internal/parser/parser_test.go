package parser

import (
	"os"
	"testing"
)

func TestValidate(t *testing.T) {
	testFiles := []struct {
		filepath string
		expected bool
	}{
		{
			filepath: "../../testdata/step1/valid.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step1/invalid.json",
			expected: false,
		},
		{
			filepath: "../../testdata/step2/valid.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step2/invalid.json",
			expected: false,
		},
		{
			filepath: "../../testdata/step2/valid2.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step2/invalid2.json",
			expected: false,
		},
		{
			filepath: "../../testdata/step3/valid_noint.json",
			expected: true,
		},
	}

	for _, testFile := range testFiles {
		file, err := os.Open(testFile.filepath)
		if err != nil {
			t.Fatalf("Error opening file %s", testFile.filepath)
		}
		defer file.Close()

		result := Validate(file)
		if result != testFile.expected {
			t.Errorf("Validation for file %s is %v, want %v", testFile.filepath, result, testFile.expected)
		}
	}
}
