package shell

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func (s *Shell) repl(){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("shell>")
		if !scanner.Scan(){
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "exit"{
			break
		}
		fmt.Println("You entered: ",line)
		s.AddHistory(line)
		s.Execute(line)
	}

} 