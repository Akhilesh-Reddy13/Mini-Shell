package shell

import "time"

type HistoryEntry struct{
	Raw string
	Time time.Time
}