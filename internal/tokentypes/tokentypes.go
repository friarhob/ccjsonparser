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
)

var names = [...]string{
	"Invalid",
	"EOF",
	"StartJSON",
	"EndJSON",
	"String",
	"Colon",
	"Comma",
}

func (t Token) String() string {
	return names[t]
}
