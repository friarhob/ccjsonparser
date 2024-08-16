package exitcodes

type ExitCodes int

const (
	ValidJSON ExitCodes = iota
	InvalidJSON
	ErrorReadingFile
	UsageError
)

var names = [...]string{
	"ValidJSON",
	"InvalidJSON",
	"ErrorReadingFile",
	"UsageError",
}

func (e ExitCodes) String() string {
	return names[e]
}
