package shell

import (
	//"os"
	"minishell/internal/parser"
)

func (s *Shell) Execute(line string) error{
	tokens := parser.Tokenize(line)
	//Remove the test printing
	/*for _,v := range tokens{
			fmt.Printf("%s - %s \n",v.Type,v.Value)
	}*/
	command,err := parser.Parse(tokens)
	if err != nil {
		return err
	}
	//fmt.Println(command)
	builtInFunc,ok  := LookUp(command.Name)
	if ok {
		err := builtInFunc(s,command)
		if err != nil {
			return err
		}
	} else {
		err := ExecuteExternal(command)
		if err != nil {
			return err
		} 
	}
	return nil
}
