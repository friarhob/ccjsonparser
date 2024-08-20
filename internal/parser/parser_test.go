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
		{
			filepath: "../../testdata/step3/valid_allnumbers.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step3/valid.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step3/invalid.json",
			expected: false,
		},
		{
			filepath: "../../testdata/step4/valid_nolist.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step4/valid2_nolist.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step4/valid.json",
			expected: true,
		},
		{
			filepath: "../../testdata/step4/invalid.json",
			expected: false,
		},
		{
			filepath: "../../testdata/step4/valid2.json",
			expected: true,
		},
		{filepath: "../../testdata/test/fail1.json", expected: false},
		{filepath: "../../testdata/test/fail2.json", expected: false},
		{filepath: "../../testdata/test/fail3.json", expected: false},
		{filepath: "../../testdata/test/fail4.json", expected: false},
		{filepath: "../../testdata/test/fail5.json", expected: false},
		{filepath: "../../testdata/test/fail6.json", expected: false},
		{filepath: "../../testdata/test/fail7.json", expected: false},
		{filepath: "../../testdata/test/fail8.json", expected: false},
		{filepath: "../../testdata/test/fail9.json", expected: false},
		{filepath: "../../testdata/test/fail10.json", expected: false},
		{filepath: "../../testdata/test/fail11.json", expected: false},
		{filepath: "../../testdata/test/fail12.json", expected: false},
		{filepath: "../../testdata/test/fail13.json", expected: false},
		{filepath: "../../testdata/test/fail14.json", expected: false},
		{filepath: "../../testdata/test/fail15.json", expected: false},
		{filepath: "../../testdata/test/fail16.json", expected: false},
		{filepath: "../../testdata/test/fail17.json", expected: false},
		{filepath: "../../testdata/test/fail18.json", expected: true}, //ignoring nesting limit
		{filepath: "../../testdata/test/fail19.json", expected: false},
		{filepath: "../../testdata/test/fail20.json", expected: false},
		{filepath: "../../testdata/test/fail21.json", expected: false},
		{filepath: "../../testdata/test/fail22.json", expected: false},
		{filepath: "../../testdata/test/fail23.json", expected: false},
		{filepath: "../../testdata/test/fail24.json", expected: false},
		{filepath: "../../testdata/test/fail25.json", expected: false},
		{filepath: "../../testdata/test/fail26.json", expected: false},
		{filepath: "../../testdata/test/fail27.json", expected: false},
		{filepath: "../../testdata/test/fail28.json", expected: false},
		{filepath: "../../testdata/test/fail29.json", expected: false},
		{filepath: "../../testdata/test/fail30.json", expected: false},
		{filepath: "../../testdata/test/fail31.json", expected: false},
		{filepath: "../../testdata/test/fail32.json", expected: false},
		{filepath: "../../testdata/test/fail33.json", expected: false},
		{filepath: "../../testdata/test/pass1.json", expected: true},
		{filepath: "../../testdata/test/pass2.json", expected: true},
		{filepath: "../../testdata/test/pass3.json", expected: true},
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
