package parser

import (
	"os"

	"github.com/friarhob/ccjsonparser/internal/lexer"
	"github.com/friarhob/ccjsonparser/internal/tokentypes"
)

func parsePair() bool {
	curToken := lexer.Consume()
	if curToken != tokentypes.String {
		return false
	}

	curToken = lexer.Consume()
	if curToken != tokentypes.Colon {
		return false
	}

	curToken = lexer.Consume()
	if curToken != tokentypes.String {
		return false
	}

	return true
}

func Validate(file *os.File) bool {
	lexer.StartLexer(file)

	curToken := lexer.Consume()
	if curToken != tokentypes.StartJSON {
		return false
	}

	hasPair := false
	hasNextPair := false

	for {
		reachEnd := false
		nextToken := lexer.Peek()
		switch nextToken {
		case tokentypes.String:
			if hasPair && !hasNextPair {
				return false
			}

			if !parsePair() {
				return false
			}

			hasPair = true
			hasNextPair = lexer.Peek() == tokentypes.Comma
			if hasNextPair {
				lexer.Consume()
			}

		case tokentypes.EndJSON:
			if hasNextPair {
				return false
			}

			lexer.Consume()
			reachEnd = true

		default:
			return false
		}

		if reachEnd {
			break
		}

	}

	curToken = lexer.Consume()
	return curToken == tokentypes.EOF
}
