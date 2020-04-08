package cli

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/urfave/cli"
	"os"
)

func add(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		var depMode = false
		if len(os.Args) < 3 {
			depMode = true
		}
		namespaces := laruche.ArrayToNamespaces(os.Args[2:])
		fmt.Println(depMode, namespaces)

		return nil
	}
}
