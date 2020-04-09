package cli

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/manager"
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

		beeOrHive, err := GetOneOfInCurrentDir()
		if err != nil {
			return err
		}

		return beeOrHive.AddDep(depMode, namespaces)
	}
}

func (bee Bee) AddDep(depMode bool, namespaces laruche.Namespaces) error {
	if depMode {
		namespaces = laruche.Bee(bee).GetSubDependencies()
	}

	var lb = laruche.Bee(bee)
	return manager.BeeAddDependencies(&lb, namespaces)
}

func (hive Hive) AddDep(depMode bool, namespaces laruche.Namespaces) error {
	if depMode {
		namespaces = laruche.Hive(hive).GetDependencies()
	}

	var lh = laruche.Hive(hive)
	return manager.HiveAddDependencies(&lh, namespaces)
}
