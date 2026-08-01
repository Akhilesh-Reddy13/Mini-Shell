package parser

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []Token
	}{
		{
			name: "single word",
			line: "pwd",
			want: []Token{{Type: TokenWord, Value: "pwd"}},
		},
		{
			name: "two words",
			line: "ls -l",
			want: []Token{
				{Type: TokenWord, Value: "ls"},
				{Type: TokenWord, Value: "-l"},
			},
		},
		{
			name: "empty input",
			line: "",
			want: []Token{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tokenize(tt.line)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Tokenize(%q) = %#v, want %#v", tt.line, got, tt.want)
			}
		})
	}
}
