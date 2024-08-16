package tokentypes

type Token int

const (
	Invalid Token = iota
	EOF
	StartJSON
	EndJSON
)

var names = [...]string{
	"Invalid",
	"EOF",
	"StartJSON",
	"EndJSON",
}

func (t Token) String() string {
	return names[t]
}
