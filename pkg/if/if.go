package _if

import (
	"fmt"
	"os"
)

func ErrorExit(title string, err error) {
	if err != nil {
		fmt.Println(title, err)
		os.Exit(0)
	}
}
