package shell

import (
)

type Shell struct{

}

func NewShell() *Shell{
	return &Shell{}
}

func (s *Shell) Run() {
	s.repl()
}