package cli

import (
	local "github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/urfave/cli"
)

func generate(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		_, err := local.GetBeeCurrentDir()
		if err != nil {
			return err
		}

		return nil
	}
}
