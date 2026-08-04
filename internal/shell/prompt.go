package shell

import (
	"os"
	"strings"
)

func (s *Shell) Prompt() string {
	cwd := s.CurrentDir
	home := os.Getenv("HOME")
	isHomePrefix := strings.HasPrefix(cwd,home)
	if isHomePrefix {
		return "~"+strings.TrimPrefix(cwd,home)
	}else{
		return cwd
	}
}