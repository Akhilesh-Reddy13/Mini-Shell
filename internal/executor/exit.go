package shell

import (
	"errors"
	"minishell/internal/parser"
	"os"
)

func Exit(s *Shell,cmd parser.Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("exit : Too many arguments.")
	}
	os.Exit(0)
	return nil
}