package cli

import (
	"github.com/benka-me/laruche/pkg/cli/scan"
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/urfave/cli"
)

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

func initHive(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		hive := laruche.InitHiveAskUser()
		setAuthor(hive)

		err := config.SaveLocal(hive)
		if err != nil {
			return err
		}
		return nil
	}
}
