package tokentypes

type Token int

const (
	Invalid Token = iota
	EOF
	StartJSON
	EndJSON
	String
	Colon
	Comma
	Boolean
	Null
)

var names = [...]string{
	"Invalid",
	"EOF",
	"StartJSON",
	"EndJSON",
	"String",
	"Colon",
	"Comma",
	"Boolean",
	"Null",
}

func (t Token) String() string {
	return names[int(t)]
}
