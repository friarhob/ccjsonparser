package lexer

import (
	"bufio"
	"errors"
	"io"
	"os"
	"unicode"

	"github.com/friarhob/ccjsonparser/internal/adt"
	"github.com/friarhob/ccjsonparser/internal/exitcodes"
	"github.com/friarhob/ccjsonparser/internal/tokentypes"
)

var reader adt.PeakableReader

var tokenBuffer *adt.Queue

func StartLexer(file *os.File) {
	reader = *adt.NewPeakableReader(bufio.NewReader(file))
	tokenBuffer = adt.NewQueue()
}

func consumeString() error {
	for {
		nextRune, err := reader.PopRune()
		if err != nil {
			return err
		}

		if nextRune == '\\' {
			_, err := reader.PopRune()
			if err != nil {
				return err
			}
		}

		if nextRune == '"' {
			return nil
		}
	}
}

func consumeReservedWord(firstRune rune) (tokentypes.Token, error) {
	undefinedErr := errors.New("reserved word undefined")
	reservedWords := map[rune]string{
		't': "rue",
		'f': "alse",
		'n': "ull",
	}

	var resToken tokentypes.Token

	_, exists := reservedWords[firstRune]
	if !exists {
		return tokentypes.Invalid, undefinedErr
	}

	if firstRune == 'n' {
		resToken = tokentypes.Null
	} else {
		resToken = tokentypes.Boolean
	}

	for _, targetRune := range reservedWords[firstRune] {
		nextRune, err := reader.PopRune()
		if err != nil {
			if err != io.EOF {
				os.Exit(int(exitcodes.ErrorReadingFile))
			}
			return tokentypes.Invalid, undefinedErr
		}
		if nextRune != targetRune {
			return tokentypes.Invalid, undefinedErr
		}
	}

	nextRune, err := reader.PeekRune()
	if err != nil {
		if err != io.EOF {
			os.Exit(int(exitcodes.ErrorReadingFile))
		}
		return resToken, nil
	}
	if unicode.IsLetter(nextRune) {
		return tokentypes.Invalid, undefinedErr
	}

	return resToken, nil
}

func generateNextToken() {
	var nextRune rune

	for {
		var err error
		nextRune, err = reader.PopRune()
		if err != nil {
			if err == io.EOF {
				tokenBuffer.Enqueue(tokentypes.EOF)
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
		tokenBuffer.Enqueue(tokentypes.StartJSON)
	case '}':
		tokenBuffer.Enqueue(tokentypes.EndJSON)
	case ':':
		tokenBuffer.Enqueue(tokentypes.Colon)
	case ',':
		tokenBuffer.Enqueue(tokentypes.Comma)
	case '"':
		err := consumeString()
		if err != nil {
			if err == io.EOF {
				tokenBuffer.Enqueue(tokentypes.Invalid)
				return
			}
			os.Exit(int(exitcodes.ErrorReadingFile))
		}
		tokenBuffer.Enqueue(tokentypes.String)
	case 't', 'f', 'n':
		resToken, err := consumeReservedWord(nextRune)
		if err != nil && err != io.EOF {
			os.Exit(int(exitcodes.ErrorReadingFile))
		}
		tokenBuffer.Enqueue(resToken)
	default:
		tokenBuffer.Enqueue(tokentypes.Invalid)
	}

}

func Peek() tokentypes.Token {
	if tokenBuffer.IsEmpty() {
		generateNextToken()
	}
	res, _ := tokenBuffer.Peek()
	return res.(tokentypes.Token)
}

func Consume() tokentypes.Token {
	if tokenBuffer.IsEmpty() {
		generateNextToken()
	}

	res, err := tokenBuffer.Dequeue()
	if err != nil {
		return tokentypes.Invalid
	}
	return res.(tokentypes.Token)
}
