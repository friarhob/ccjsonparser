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
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step2/invalid.json",
			expected: []tokentypes.Token{tokentypes.StartJSON, tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step3/valid_noint.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.Boolean, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Boolean, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Null, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step3/valid.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.Boolean, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Boolean, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Null, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step3/valid_allnumbers.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step3/valid_allnumbers.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step4/valid_nolist.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.StartJSON, tokentypes.EndJSON,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step4/valid2_nolist.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.String,
				tokentypes.EndJSON,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step4/valid.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.StartJSON, tokentypes.EndJSON, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.StartList, tokentypes.EndList,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/step4/valid2.json",
			expected: []tokentypes.Token{tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.String,
				tokentypes.EndJSON, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.StartList, tokentypes.String, tokentypes.EndList,
				tokentypes.EndJSON, tokentypes.EOF},
		},
		{
			filepath: "../../testdata/test/fail1.json",
			expected: []tokentypes.Token{tokentypes.String,
				tokentypes.EOF},
		},
		{
			filepath: "../../testdata/test/pass1.json",
			expected: []tokentypes.Token{tokentypes.StartList,
				tokentypes.String, tokentypes.Comma,
				tokentypes.StartJSON, tokentypes.String, tokentypes.Colon, tokentypes.StartList, tokentypes.String, tokentypes.EndList, tokentypes.EndJSON, tokentypes.Comma,
				tokentypes.StartJSON, tokentypes.EndJSON, tokentypes.Comma,
				tokentypes.StartList, tokentypes.EndList, tokentypes.Comma,
				tokentypes.Number, tokentypes.Comma,
				tokentypes.Boolean, tokentypes.Comma, //21
				tokentypes.Boolean, tokentypes.Comma,
				tokentypes.Null, tokentypes.Comma,
				tokentypes.StartJSON,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma, //30
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma, // 50
				tokentypes.String, tokentypes.Colon, tokentypes.Number, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, //70
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, //90
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Boolean, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Boolean, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.Null, tokentypes.Comma, //110
				tokentypes.String, tokentypes.Colon, tokentypes.StartList, tokentypes.EndList, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.StartJSON, tokentypes.EndJSON, tokentypes.Comma, //120
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma,
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, //136
				tokentypes.String, tokentypes.Colon, tokentypes.StartList, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number,

				tokentypes.Comma, //145

				tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.EndList, tokentypes.Comma, tokentypes.String, tokentypes.Colon, tokentypes.StartList, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.EndList, tokentypes.Comma, //172
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, // 176
				tokentypes.String, tokentypes.Colon, tokentypes.String, tokentypes.Comma, // 180
				tokentypes.String,
				tokentypes.Colon, tokentypes.String,
				tokentypes.EndJSON, tokentypes.Comma,
				tokentypes.Number, tokentypes.Comma, tokentypes.Number,
				tokentypes.Comma,
				tokentypes.Number, // 190
				tokentypes.Comma,

				tokentypes.Number, tokentypes.Comma,
				tokentypes.Number, tokentypes.Comma,
				tokentypes.Number, tokentypes.Comma,
				tokentypes.Number, tokentypes.Comma,
				tokentypes.Number, tokentypes.Comma, tokentypes.Number, tokentypes.Comma, tokentypes.Number,
				tokentypes.Comma, tokentypes.String, tokentypes.EndList, tokentypes.EOF},
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
			t := Consume()
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
