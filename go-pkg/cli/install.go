package cli

import (
	"errors"
	"github.com/urfave/cli"
	"os"
)

func install(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		if len(os.Args) < 3 {
			return errors.New("usage: laruche install [author/name]")
		}
		return nil
	}
}
