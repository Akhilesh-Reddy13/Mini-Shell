package main

import (
	"minishell/internal/shell"
)


func main(){
	shell := shell.NewShell()

	shell.Run()
	
}