package shell

import (
	"log"
	"os"
	"strings"
)

type Shell struct{
	CurrentDir string
	Env map[string]string
	History []HistoryEntry
	LastExitCode int
}

func NewShell() *Shell{
	cwd,err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	envString := os.Environ()
	env := make(map[string]string)
	for _,v := range envString{
		result :=strings.SplitN(v,"=",2)
		if len(result) != 2{
			continue
		}
		env[result[0]] = result[1]
	}
	return &Shell{
		CurrentDir: cwd,
		Env: env,
		History: []HistoryEntry{},
		LastExitCode: 0,
	}
}

func (s *Shell) Run() {
	s.Repl()
}