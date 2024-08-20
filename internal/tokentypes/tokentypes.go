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
	StartList
	EndList
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
	"StartList",
	"EndList",
}

func (t Token) String() string {
	return names[int(t)]
}
