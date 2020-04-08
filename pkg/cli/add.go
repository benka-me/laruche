package cli

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/local"
	"github.com/urfave/cli"
	"os"
)

func add(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		var depMode = false
		if len(os.Args) < 3 {
			depMode = true
		}
		namespaces, err := laruche.ArrayToNamespaces(os.Args[2:])
		if err != nil {
			return err
		}

		beeOrHive, err := local.GetOneOfCurrentDir()
		if err != nil {
			return err
		}

		return beeOrHive.AddDep(depMode, namespaces)
	}
}
