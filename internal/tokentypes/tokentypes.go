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
	Number
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
	"Number",
}

func (t Token) String() string {
	return names[int(t)]
}
