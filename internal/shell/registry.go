package shell

import (
	"minishell/internal/parser"
)

type BuiltinFunc func(*Shell, parser.Command) error 

var registry = map[string]BuiltinFunc{
	/*"cd" : ,
	"exit" : ,
	"history" : , 
	"pwd" :*/ 
}

func LookUp(cmd string) (BuiltinFunc, bool) {
	value,ok := registry[cmd]
	if !ok{
		return nil,false
	}
	return value,true
}