package shell

import (
	"fmt"
	//"os"
	"minishell/internal/parser"
)

func (s *Shell) Execute(line string) {
	tokens := parser.Tokenize(line)
	//Remove the test printing
	/*for _,v := range tokens{
			fmt.Printf("%s - %s \n",v.Type,v.Value)
	}*/
	command,err := parser.Parse(tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(command)
	builtInFunc,ok  := LookUp(command.Name)
	if ok {
		err := builtInFunc(s,command)
		if err != nil {
			fmt.Println(err)
		}
	}
	if !ok {
		fmt.Println("builtin not found:", command.Name)
	}
	//return executeExternal(command)
	
}
