package lexer

import (
	"bufio"
	"io"
	"os"
	"unicode"

	"github.com/friarhob/ccjsonparser/internal/adt"
	"github.com/friarhob/ccjsonparser/internal/exitcodes"
	"github.com/friarhob/ccjsonparser/internal/tokentypes"
)

var reader bufio.Reader

var buffer *adt.Queue

func StartLexer(file *os.File) {
	reader = *bufio.NewReader(file)
	buffer = adt.NewQueue()
}

func generateNextToken() {
	var nextRune rune

	for {
		var err error
		nextRune, _, err = reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				buffer.Enqueue(tokentypes.EOF)
				return
			}

			os.Exit(int(exitcodes.ErrorReadingFile))
		}

		if !unicode.IsSpace(nextRune) {
			break
		}
	}

	switch nextRune {
	case '{':
		buffer.Enqueue(tokentypes.StartJSON)
	case '}':
		buffer.Enqueue(tokentypes.EndJSON)
	default:
		buffer.Enqueue(tokentypes.Invalid)
	}

}

func Peek() tokentypes.Token {
	if buffer.IsEmpty() {
		generateNextToken()
	}
	res, _ := buffer.Peek()
	return res.(tokentypes.Token)
}

func Consume() tokentypes.Token {
	if buffer.IsEmpty() {
		generateNextToken()
	}

	res, _ := buffer.Dequeue()
	return res.(tokentypes.Token)
}
