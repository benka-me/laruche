package cli

import (
	"github.com/benka-me/laruche/go-pkg/cli/form"
	"github.com/benka-me/laruche/go-pkg/cli/scan"
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/benka-me/laruche/go-pkg/generator"
	"github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"github.com/urfave/cli"
)

func initGateway(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := form.InitGatewayAskUser()
		return ActionInitGateway(bee)
	}
}
func initClient(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := form.InitClientAskUser()
		return ActionInitClient(bee)
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

	err = generator.GenerateMainService(bee)
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

	err = generator.GenerateMainService(bee)
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

	err = generator.GenerateMainClient(bee)
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
