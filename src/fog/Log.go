package fog

import (
	"fmt"
)

func WriteLog(text string) {
	fmt.Println(text)
}

func WriteLogResult(e error) {
	if e != nil {
		WriteLog(e.Error())
	}
}
