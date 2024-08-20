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
			escapedRune, err := reader.PopRune()
			if err != nil {
				return err
			}

			invalidError := errors.New("invalid escaped rune")

			switch escapedRune {
			case '"', '\\', '/', 'b', 'f', 'n', 'r':
			case 'u':
				for i := 0; i < 4; i++ {
					unicodeRune, err2 := reader.PopRune()
					if err2 != nil {
						return err
					}
					if !unicode.Is(unicode.Hex_Digit, unicodeRune) {
						return invalidError
					}
				}
			default:
				return invalidError
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

func consumeNumber(firstRune rune) error {
	type state int
	const (
		stateInvalid state = iota
		stateNegative
		stateInteger
		stateZero
		stateDecimal
		stateDecimalValue
		stateExponential
		stateExponentialSignal
		stateExponentialValue
	)

	var curState state
	invalidError := errors.New("invalid number")

	switch firstRune {
	case '-':
		curState = stateNegative
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		curState = stateInteger
	case '0':
		curState = stateZero
	default:
		curState = stateInvalid
	}

	for {
		switch curState {

		case stateInvalid:
			return invalidError

		case stateNegative:
			nextRune, err := reader.PopRune()
			if err != nil {
				if err == io.EOF {
					err = invalidError
				}
				return invalidError
			}

			switch nextRune {
			case '.':
				curState = stateDecimal
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				curState = stateInteger
			case '0':
				curState = stateZero
			default:
				curState = stateInvalid
			}

		case stateZero:
			nextRune, err := reader.PeekRune()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return invalidError
			}

			if unicode.IsSpace(nextRune) || nextRune == ',' || nextRune == '}' || nextRune == '[' {
				return nil
			}

			nextRune, err = reader.PopRune()
			if err != nil {
				os.Exit(int(exitcodes.ErrorReadingFile))
			}

			switch nextRune {
			case '.':
				curState = stateDecimal
			case 'e', 'E':
				curState = stateExponential
			default:
				curState = stateInvalid
			}

		case stateInteger:
			nextRune, err := reader.PeekRune()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return invalidError
			}

			if unicode.IsSpace(nextRune) || nextRune == ',' || nextRune == '}' || nextRune == '[' {
				return nil
			}

			nextRune, err = reader.PopRune()
			if err != nil {
				os.Exit(int(exitcodes.ErrorReadingFile))
			}

			switch nextRune {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			case '.':
				curState = stateDecimal
			case 'e', 'E':
				curState = stateExponential
			default:
				curState = stateInvalid
			}

		case stateDecimal:
			nextRune, err := reader.PopRune()
			if err != nil {
				if err == io.EOF {
					err = invalidError
				}
				return invalidError
			}

			switch nextRune {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				curState = stateDecimalValue
			default:
				curState = stateInvalid
			}

		case stateDecimalValue:
			nextRune, err := reader.PeekRune()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return invalidError
			}

			if unicode.IsSpace(nextRune) || nextRune == ',' || nextRune == '}' || nextRune == '[' {
				return nil
			}

			nextRune, err = reader.PopRune()
			if err != nil {
				os.Exit(int(exitcodes.ErrorReadingFile))
			}

			switch nextRune {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			case 'e', 'E':
				curState = stateExponential
			default:
				curState = stateInvalid

			}

		case stateExponential:
			nextRune, err := reader.PopRune()
			if err != nil {
				if err == io.EOF {
					err = invalidError
				}
				return invalidError
			}

			switch nextRune {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				curState = stateExponentialValue
			case '+', '-':
				curState = stateExponentialSignal
			default:
				curState = stateInvalid
			}

		case stateExponentialSignal:
			nextRune, err := reader.PopRune()
			if err != nil {
				if err == io.EOF {
					err = invalidError
				}
				return invalidError
			}

			switch nextRune {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				curState = stateExponentialValue
			default:
				curState = stateInvalid
			}

		case stateExponentialValue:
			nextRune, err := reader.PeekRune()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return invalidError
			}

			if unicode.IsSpace(nextRune) || nextRune == ',' || nextRune == '}' || nextRune == '[' {
				return nil
			}

			nextRune, err = reader.PopRune()
			if err != nil {
				os.Exit(int(exitcodes.ErrorReadingFile))
			}

			switch nextRune {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				curState = stateInvalid

			}
		}
	}
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
	case '[':
		tokenBuffer.Enqueue(tokentypes.StartList)
	case ']':
		tokenBuffer.Enqueue(tokentypes.EndList)
	case ':':
		tokenBuffer.Enqueue(tokentypes.Colon)
	case ',':
		tokenBuffer.Enqueue(tokentypes.Comma)
	case '"':
		err := consumeString()
		if err != nil {
			if err == io.EOF || err.Error() == "invalid escaped rune" {
				tokenBuffer.Enqueue(tokentypes.Invalid)
				return
			}
			os.Exit(int(exitcodes.ErrorReadingFile))
		}
		tokenBuffer.Enqueue(tokentypes.String)
	case 't', 'f', 'n':
		resToken, err := consumeReservedWord(nextRune)
		if err != nil && err != io.EOF && err.Error() != "reserved word undefined" {
			os.Exit(int(exitcodes.ErrorReadingFile))
		}
		tokenBuffer.Enqueue(resToken)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		err := consumeNumber(nextRune)
		if err != nil {
			if err == io.EOF || err.Error() == "invalid number" {
				tokenBuffer.Enqueue(tokentypes.Invalid)
				return
			}
			os.Exit(int(exitcodes.ErrorReadingFile))
		}
		tokenBuffer.Enqueue(tokentypes.Number)
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
