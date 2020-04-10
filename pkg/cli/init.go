package cli

import (
	"github.com/benka-me/laruche/pkg/cli/form"
	"github.com/benka-me/laruche/pkg/cli/scan"
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/generator"
	"github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/urfave/cli"
)

func ActionInitBee(bee *laruche.Bee) error {
	err := config.AddBee(bee)
	if err != nil {
		return nil
	}

	err = generator.GenerateAll(bee)
	if err != nil {
		return nil
	}

	return local.SaveBee(bee)
}
func initBee(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := form.InitBeeAskUser()
		return ActionInitBee(bee)
	}
}
func ActionInitHive(hive *laruche.Hive) error {
	setAuthor(hive)

	err := local.SaveHive(hive)
	if err != nil {
		return err
	}
	return nil
}

func initHive(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		hive := laruche.InitHiveAskUser()
		return ActionInitHive(hive)
	}
}
func setAuthor(hive *laruche.Hive) {
	author := config.GetState().Username
	if author == "" {
		hive.Name = scan.Step(
			"No username found, please login or type a temporary username (not recommended):",
			"required,lte=20,gte=2",
			func(s string) error { return nil })
	} else {
		hive.Author = author
	}
}
