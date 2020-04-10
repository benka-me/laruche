package cli

import (
	local2 "github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/urfave/cli"
)

func generate(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		_, err := local2.GetBeeCurrentDir()
		if err != nil {
			return err
		}

		return nil
	}
}
