package ls

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func SimpleLs() {
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

		lsout.inode, errParse = strconv.Atoi(temp[0])
		if errParse != nil {
			panic("unable to parse inode value")
		}

		lsout.permission = temp[1]
		lsout.no_of_content, errParse = strconv.Atoi(temp[2])
		if errParse != nil {
			panic("unable to parse no_content integer")
		}

		lsout.user = temp[3]
		lsout.group = temp[4]
		lsout.author = temp[5]

		lsout.size, errParse = strconv.Atoi(temp[6])
		if errParse != nil {
			panic("unable to parse size integer")
		}

		lsout.filename = temp[10]
		timeString := temp[7] + "T" + temp[8] + temp[9][:3] + ":" + temp[9][3:]

		lsout.datetime, err = time.Parse(time.RFC3339Nano, timeString)
		if err != nil {
			fmt.Println("unable to parse time integer : ", err)
		}

		lsOutput = append(lsOutput, lsout)
	}

	fmt.Printf("%+v", lsOutput)
}
