package cli

import (
	"github.com/urfave/cli"
)

func Hive(app App) cli.ActionFunc {
	return func(context *cli.Context) error {
		return nil
	}
}
