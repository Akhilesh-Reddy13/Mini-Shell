package parser

type TokenType int

const (
	TokenWord TokenType = iota
	TokenPipe
	TokenRedirectIn
	TokenRedirectOut
)

type Token struct {
	Type  TokenType
	Value string
}

func (t TokenType) String() string{
	switch t{
	case TokenWord:
		return "Word"
	case TokenPipe:
		return  "Pipe"
	case TokenRedirectIn:
		return "RedirectIn"
	case TokenRedirectOut:
		return "RedirectOut"
	default:
		return "UNKNOWN"
	}
}