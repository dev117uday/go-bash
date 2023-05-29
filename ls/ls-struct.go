package ls

import "time"

type LsStruct struct {
	inode         int
	permission    string
	no_of_content int
	user          string
	group         string
	author        string
	size          int
	datetime      time.Time
	filename      string
}
