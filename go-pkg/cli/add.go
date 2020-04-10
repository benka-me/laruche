package cli

import (
	"github.com/benka-me/laruche/go-pkg/laruche"
	"github.com/benka-me/laruche/go-pkg/manager"
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

		err = beeOrHive.AddDep(depMode, namespaces)
		if err != nil {
			return err
		}

		return nil
	}
}

func (bee Bee) AddDep(depMode bool, namespaces laruche.Namespaces) error {
	if depMode {
		namespaces = laruche.Bee(bee).GetDependencies()
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
