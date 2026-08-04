package shell

import (
	"errors"
	"fmt"
	"minishell/internal/parser"
)

func WorkDir(s *Shell,cmd parser.Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("pwd : Too many arguments.")
	}
	fmt.Println(s.CurrentDir)
	return nil
}