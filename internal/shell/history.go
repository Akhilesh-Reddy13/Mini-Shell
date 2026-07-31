package shell

import (
	"fmt"
	"time"
)

type HistoryEntry struct{
	Raw string
	Time time.Time
}

func (s *Shell) AddHistory(line string) {
	entry := HistoryEntry{
		Raw: line,
		Time: time.Now(),
	}
	s.History = append(s.History,entry)
}

func (s *Shell) PrintHistory(){
	for i,v := range s.History{
		fmt.Printf("%d\t%s\n",(i+1),v.Raw)
	}
}
