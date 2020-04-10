package cli

import (
	"github.com/benka-me/laruche/go-pkg/generator"
	local2 "github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/urfave/cli"
)

func protoc(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee, err := local2.GetBeeCurrentDir()
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
