package cli

import (
	"github.com/benka-me/laruche/pkg/config"
	"github.com/benka-me/laruche/pkg/generator"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/urfave/cli"
)

func initBee(app App) cli.ActionFunc {
	return func(context *cli.Context) error {
		bee := laruche.InitBeeAskUser()
		config.AddBee(bee)
		generator.GenerateAll(bee)
		return bee.SaveLocal(config.SourcePath)
	}
}
