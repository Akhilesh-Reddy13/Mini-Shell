package shell

import "fmt"

func (s *Shell) Execute(line string) {
	if line == "history"{
		s.PrintHistory()
		return
	}
	fmt.Println("Unknown Command")
}