package shell

import (
	"minishell/internal/parser"
	"os/exec"
	"os"
)


func ExecuteExternal(cmd parser.Command) error {
	cmdName := cmd.Name

	execCmd := exec.Command(cmdName, cmd.Args...)
	// Bind the parent process standard streams to the child process
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err := execCmd.Run()
	if err != nil {
		return err
	}
	return nil
}