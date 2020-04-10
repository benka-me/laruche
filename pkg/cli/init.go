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

func initGateway(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := form.InitServiceAskUser()
		return ActionInitService(bee)
	}
}
func initClient(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := form.InitServiceAskUser()
		return ActionInitService(bee)
	}
}
func initService(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := form.InitServiceAskUser()
		return ActionInitService(bee)
	}
}
func initHive(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		hive := laruche.InitHiveAskUser()
		return ActionInitHive(hive)
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

func ActionInitService(bee *laruche.Bee) error {
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

func ActionInitGateway(bee *laruche.Bee) error {
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
func ActionInitClient(bee *laruche.Bee) error {
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
