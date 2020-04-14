package cli

import (
	"fmt"
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/benka-me/laruche/go-pkg/discover"
	"github.com/benka-me/laruche/go-pkg/http/rpc"
	urfave "github.com/urfave/cli"
	"google.golang.org/grpc"
	"os"
	"sort"
)

type App struct {
	rpc.Clients
	State *config.State
}

func Run() {
	engine, err := discover.ParseEngine("benka-me/laruche-hive", true)
	if err != nil {
		fmt.Println(err)
		return
	}
	app := &App{
		Clients: rpc.InitClients(*engine, grpc.WithInsecure()), // Init clients of dependencies services, please change options.
		State:   config.Init(),
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
					Name:   "service",
					Action: initService(app),
					Usage:  "init bee (micro-service)",
				},
				{
					Name:   "gateway",
					Action: initGateway(app),
					Usage:  "init bee (micro-service)",
				},
				{
					Name:   "client",
					Action: initClient(app),
					Usage:  "init bee (micro-service)",
				},
				{
					Name:   "hive",
					Action: initHive(app),
					Usage:  "init hive (server)",
				},
			},
		},
		{
			Name:   "add",
			Action: add(app),
			Usage:  "add dependency to current hive or bee",
		},
		{
			Name:   "protoc",
			Action: protoc(app),
			Usage:  "generate protobuf code",
		},
		{
			Name:   "generate",
			Usage:  "generate files",
			Action: generate(app),
		},
		{
			Name:   "publish",
			Action: publish(app),
			Usage:  "publish on hive-and-bees.com",
		},
		{
			Name:   "privatize",
			Action: privatize(app),
			Usage:  "privatish on hive-and-bees.com",
		},
		{
			Name:   "push",
			Action: push(app),
			Usage:  "push on hive-bees.com",
		},
		{
			Name:   "register",
			Action: register(app),
			Usage:  "register to hive-micro-bee.com",
		},
		{
			Name:   "login",
			Action: login(app),
			Usage:  "login",
		},
		{
			Name:   "whoami",
			Action: whoAmI(app),
			Usage:  "Don't forget who you are",
		},
	}
	cliApp.Action = func(context *urfave.Context) error {
		return nil
	}

	sort.Sort(urfave.FlagsByName(cliApp.Flags))
	sort.Sort(urfave.CommandsByName(cliApp.Commands))

	err = cliApp.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
