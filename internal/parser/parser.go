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

	switch lexer.Peek() {
	case tokentypes.StartJSON:
		if !parseObject() {
			return false
		}
		return true

	case tokentypes.StartList:
		if !parseList() {
			return false
		}
		return true

	case tokentypes.String, tokentypes.Boolean, tokentypes.Null, tokentypes.Number:
		_ = lexer.Consume()

		return true

	default:
		return false
	}
}

func parseObject() bool {
	curToken := lexer.Consume()
	if curToken != tokentypes.StartJSON {
		return false
	}

	hasPair := false
	hasNextPair := false

	for {
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
			return true

		default:
			return false
		}
	}
}

func parseList() bool {
	curToken := lexer.Consume()
	if curToken != tokentypes.StartList {
		return false
	}

	hasElem := false
	hasNextElem := false

	for {
		nextToken := lexer.Peek()
		switch nextToken {
		case tokentypes.String, tokentypes.Boolean, tokentypes.Null, tokentypes.Number:
			if hasElem && !hasNextElem {
				return false
			}
			lexer.Consume()

			hasElem = true
			hasNextElem = lexer.Peek() == tokentypes.Comma
			if hasNextElem {
				lexer.Consume()
			}

		case tokentypes.EndList:
			if hasNextElem {
				return false
			}

			lexer.Consume()
			return true

		case tokentypes.StartJSON:
			if hasElem && !hasNextElem {
				return false
			}

			if !parseObject() {
				return false
			}

			hasElem = true
			hasNextElem = lexer.Peek() == tokentypes.Comma
			if hasNextElem {
				lexer.Consume()
			}

		case tokentypes.StartList:
			if hasElem && !hasNextElem {
				return false
			}

			if !parseList() {
				return false
			}

			hasElem = true
			hasNextElem = lexer.Peek() == tokentypes.Comma
			if hasNextElem {
				lexer.Consume()
			}

		default:
			return false
		}
	}
}

func Validate(file *os.File) bool {
	lexer.StartLexer(file)

	switch lexer.Peek() {
	case tokentypes.StartJSON:
		if !parseObject() {
			return false
		}

		curToken := lexer.Consume()
		return curToken == tokentypes.EOF
	default:
		return false
	}

}
