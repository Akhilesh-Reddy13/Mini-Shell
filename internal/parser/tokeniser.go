package parser

import (
	"unicode"
)

func Tokenize(line string) []Token {
	runes := []rune(line)
	tokens := []Token{}

	cursor := 0

	for cursor < len(runes) {

		
		for cursor < len(runes) && unicode.IsSpace(runes[cursor]) {
			cursor++
		}

		
		if cursor >= len(runes) {
			break
		}

	
		if runes[cursor] == '|' {
			tokens = append(tokens, Token{
				Type:  TokenPipe,
				Value: "|",
			})
			cursor++
			continue
		}

		
		if runes[cursor] == '<' {
			tokens = append(tokens, Token{
				Type:  TokenRedirectIn,
				Value: "<",
			})
			cursor++
			continue
		}

		
		if runes[cursor] == '>' {
			tokens = append(tokens, Token{
				Type:  TokenRedirectOut,
				Value: ">",
			})
			cursor++
			continue
		}

		
		start := cursor

		for cursor < len(runes) &&
			!unicode.IsSpace(runes[cursor]) &&
			runes[cursor] != '|' &&
			runes[cursor] != '<' &&
			runes[cursor] != '>' {

			cursor++
		}

		tokens = append(tokens, Token{
			Type:  TokenWord,
			Value: string(runes[start:cursor]),
		})
	}

	return tokens
}