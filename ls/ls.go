package ls

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetLs() []LsStruct {
	cmd := exec.Command("ls", "-A", "-lt", "--author", "-b", "-c", "-i", "--time-style=full-iso")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	output_arr := strings.Split(string(stdout), "\n")
	output_arr = output_arr[1 : len(output_arr)-1]

	var lsOutput []LsStruct

	for _, data := range output_arr {

		data = strings.Join(strings.Fields(strings.TrimSpace(data)), " ")
		temp := strings.Split(data, " ")

		var lsout LsStruct
		var errParse error

		lsout.Inode, errParse = strconv.Atoi(temp[0])
		if errParse != nil {
			panic("unable to parse inode value")
		}

		lsout.Permission = temp[1]
		lsout.No_of_content, errParse = strconv.Atoi(temp[2])
		if errParse != nil {
			panic("unable to parse no_content integer")
		}

		lsout.User = temp[3]
		lsout.Group = temp[4]
		lsout.Author = temp[5]

		lsout.Size, errParse = strconv.Atoi(temp[6])
		if errParse != nil {
			panic("unable to parse size integer")
		}

		lsout.Filename = temp[10]
		timeString := temp[7] + "T" + temp[8] + temp[9][:3] + ":" + temp[9][3:]

		lsout.Datetime, err = time.Parse(time.RFC3339Nano, timeString)
		if err != nil {
			fmt.Println("unable to parse time integer : ", err)
		}

		lsOutput = append(lsOutput, lsout)
	}

	// fmt.Printf("%+v", lsOutput)
	return lsOutput
}
