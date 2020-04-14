package scan

import (
	"bufio"
	"fmt"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

var V *validator.Validate

func fmtError(err string) string {
	switch err {
	case "required":
		return "Empty field is required"
	case "gte":
		return "Too short"
	case "lte":
		return "Too long"
	case "alpha":
		return "Alphabetic characters only"
	case "number":
		return "please enter valid number"
	case "max":
		return "too high"
	case "min":
		return "too low"
	case "email":
		return "please enter valid email"
	default:
		return ""
	}
}

func GetPasswordTwice(title string) string {
	fmt.Printf("%s : ", title)
	password, err := terminal.ReadPassword(0)

	if err != nil {
		fmt.Println(err)
		return GetPassword(title)
	}
	fmt.Printf("\n%s : ", "Confirm password")
	confirm, err := terminal.ReadPassword(0)

	if string(password) != string(confirm) {
		fmt.Println("passwords don't match\n-----------------------------------------")
		return GetPasswordTwice(title)
	}

	return string(password)
}

func GetPassword(title string) string {
	fmt.Printf("%s : ", title)
	password, err := terminal.ReadPassword(0)

	if err != nil {
		fmt.Println(err)
		return GetPassword(title)
	}

	fmt.Println("")
	return string(password)
}

func StepBool(title string) bool {
	var tmp string

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n): ", title)
	wn, _ := reader.ReadString('\n')
	tmp = strings.Replace(wn, "\n", "", 1)
	if tmp != "y" && tmp != "n" {
		fmt.Println("y or n only, please try again")
		return StepBool(title)
	}

	return tmp == "y"
}
func NoFunc(s string) error {
	return nil
}
func Step(title, validate string, custom StringValid) string {
	var tmp string

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s : ", title)
	tmp, _ = reader.ReadString('\n')
	tmp = strings.Replace(tmp, "\n", "", -1)
	fmt.Println(tmp)

	err := custom(tmp)
	if err != nil {
		fmt.Println(err)
		return Step(title, validate, custom)
	} else if err := V.Var(tmp, validate); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(fmtError(e.Tag()))
		}
		return Step(title, validate, custom)
	}
	return tmp
}
