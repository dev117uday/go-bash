package ls

import "time"

type LsStruct struct {
	Inode         int
	Permission    string
	No_of_content int
	User          string
	Group         string
	Author        string
	Size          int
	Datetime      time.Time
	Filename      string
}
