package cli

import (
	"github.com/benka-me/laruche/pkg/config"
	urfave "github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

type App struct {
	config.State
}

func Run() {
	app := App{}
	cliApp := urfave.NewApp()
	cliApp.Name = "Hive"
	cliApp.Usage = "Manage your micro-services based server"

	cliApp.Flags = []urfave.Flag{
		&urfave.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		&urfave.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
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
				//{
				//	Name:   "hive",
				//	Action: initier.Hive,
				//	Usage:  "init hive application",
				//},
			},
		},
	}
	cliApp.Action = func(context *urfave.Context) error {
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
