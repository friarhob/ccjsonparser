package lexer

import (
	"os"
	"reflect"
	"testing"

	"github.com/friarhob/ccjsonparser/internal/tokentypes"
)

func TestConsume(t *testing.T) {
	testFiles := []struct {
		filepath string
		expected []tokentypes.Token
	}{
		{
			filepath: "../../testdata/step1/valid.json",
			expected: []tokentypes.Token{tokentypes.StartJSON, tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step2/valid.json",
			expected: []tokentypes.Token{tokentypes.StartJSON, tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step2/valid2.json",
			expected: []tokentypes.Token{tokentypes.StartJSON, tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step2/invalid.json",
			expected: []tokentypes.Token{tokentypes.StartJSON, tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, tokentypes.EndJSON, tokentypes.EOF},
		},
	}

	for _, testFile := range testFiles {
		file, err := os.Open(testFile.filepath)
		if err != nil {
			t.Fatalf("Error opening file %s", testFile.filepath)
		}
		defer file.Close()
		StartLexer(file)

		result := []tokentypes.Token{}
		for {
			var t tokentypes.Token
			t = Consume()
			result = append(result, t)
			if t == tokentypes.EOF {
				break
			}
		}

		if !reflect.DeepEqual(result, testFile.expected) {
			t.Errorf("Validation for file %s is %v, want %v", testFile.filepath, result, testFile.expected)
		}
	}
}
