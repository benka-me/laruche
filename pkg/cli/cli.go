package cli

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/config"
	urfave "github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

type App struct {
	State *config.State
}

func Run() {
	app := App{
		State: config.Init(),
	}
	cliApp := urfave.NewApp()
	cliApp.Name = "Laruche"
	cliApp.Usage = "Manage your micro-services based server"

	cliApp.Commands = urfave.Commands{
		{
			Name:  "init",
			Usage: "init bee or hive",
			Subcommands: urfave.Commands{
				{
					Name:   "bee",
					Action: initBee(app),
					Usage:  "init bee (micro-service)",
				},
			},
		},
	}
	cliApp.Action = func(context *urfave.Context) error {
		fmt.Println(config.GetBee("benka/test").GetPath())
		return nil
	}

	sort.Sort(urfave.FlagsByName(cliApp.Flags))
	sort.Sort(urfave.CommandsByName(cliApp.Commands))

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}
