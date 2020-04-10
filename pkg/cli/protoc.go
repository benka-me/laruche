package cli

import (
	"github.com/benka-me/laruche/pkg/generator"
	local "github.com/benka-me/laruche/pkg/get-local"
	"github.com/urfave/cli"
)

func protoc(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee, err := local.GetBeeCurrentDir()
		if err != nil {
			return err
		}

		err = generator.Protoc(bee)
		if err != nil {
			return err
		}
		return nil
	}
}
