package parser

import (
	"os"

	"github.com/friarhob/ccjsonparser/internal/lexer"
	"github.com/friarhob/ccjsonparser/internal/tokentypes"
)

func Validate(file *os.File) bool {
	lexer.StartLexer(file)

	curToken := lexer.Consume()
	if curToken != tokentypes.StartJSON {
		return false
	}

	curToken = lexer.Consume()
	if curToken != tokentypes.EndJSON {
		return false
	}

	curToken = lexer.Consume()
	return curToken == tokentypes.EOF
}
