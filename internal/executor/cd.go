package shell

import (
	"errors"
	"minishell/internal/parser"
	"os"
)

func ChangeDirectory(s *Shell,cmd parser.Command) error{
	
	if len(cmd.Args)>1 {
		return errors.New("cd : Too many arguments.")
	}
	var path string
	if len(cmd.Args) == 0 {
		path = os.Getenv("HOME")
		if path == ""{
			return errors.New("cd : HOME not set.")
		}

	} else {
	path = cmd.Args[0]
	}
	
	if err := os.Chdir(path); err != nil{
		return err
	}

	cwd,er := os.Getwd()
	if er != nil{
			return er
	}
	s.CurrentDir = cwd

	return nil
}