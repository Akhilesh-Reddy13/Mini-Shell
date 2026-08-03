package parser

import "errors"

type Command struct {
	Name string
	Args []string
}

func Parse(tokens []Token) (Command, error) {

	if len(tokens) == 0{
		return Command{},errors.New("No tokens found")
	}

	if tokens[0].Type != TokenWord{
		return Command{},errors.New("No Command found")
	}
	cmd := Command{
		Name : tokens[0].Value,
	} 
	//Current grammar is Command -> Word Word*
	for _,v := range tokens[1:] {
		if v.Type != TokenWord {
			return Command{},errors.New("Current Grammar does not non-words")
		}
		cmd.Args = append(cmd.Args, v.Value) 
	}
	return cmd,nil
}