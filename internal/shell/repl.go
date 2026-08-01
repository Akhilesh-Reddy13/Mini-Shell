package shell

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func (s *Shell) Repl(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(s.Prompt)
		if !scanner.Scan(){
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		s.AddHistory(line)
		s.Execute(line)
	}

} 