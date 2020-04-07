package cli

import (
	"fmt"
	"github.com/benka-me/hive-server-core/go-pkg/core"
	"github.com/urfave/cli"
)

func initBee(app App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := core.Bee{}

		fmt.Println(bee)
		return nil
	}
}
