package cli

import (
	"github.com/benka-me/laruche/pkg/cli/form"
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/generator"
	"github.com/benka-me/laruche/pkg/local"
	"github.com/urfave/cli"
)

func initBee(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := form.InitBeeAskUser()
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
}
