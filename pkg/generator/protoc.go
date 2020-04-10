package generator

import (
	"bytes"
	"fmt"
	"os/exec"
)

func runProtocCommand(args []string) {
	cmd := exec.Command("protoc", args...)
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errs
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(errs.String())
	}
	fmt.Printf("%s\n", out.String())
}
