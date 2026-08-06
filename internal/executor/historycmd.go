package shell

import (
	"errors"
	"minishell/internal/parser"
)

func History(s *Shell,cmd parser.Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("exit : Too many arguments.")
	}
	s.PrintHistory()
	return nil
}