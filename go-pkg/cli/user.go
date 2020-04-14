package cli

import (
	"context"
	"fmt"
	"github.com/benka-me/laruche/go-pkg/cli/scan"
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/benka-me/users/go-pkg/users"
	"github.com/go-playground/validator"
	"github.com/urfave/cli"
)

func login(app *App) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		req := &users.LoginReq{}

		username := scan.Step(
			"Username",
			"required,lte=20,gte=3",
			func(s string) error { return nil })

		req.Identifier = username
		req.Password = scan.GetPassword("Password")

		_, err := app.Users.Login(context.TODO(), req)
		if err != nil {
			return err
		}
		fmt.Println("Success")
		return nil
	}
}
func register(app *App) cli.ActionFunc {
	scan.V = validator.New()
	return func(ctx *cli.Context) error {
		req := &users.RegisterReq{}
		req.Username = scan.Step("Username",
			"required,lte=20,gte=3",
			scan.NoFunc)
		req.Email = scan.Step("Email", "email", scan.NoFunc)
		req.Password = scan.GetPasswordTwice("Password")

		_, err := app.UsersGateway.Register(context.TODO(), req)
		if err != nil {
			return err
		}
		fmt.Println("Success")
		return nil
	}
}
func whoAmI(app *App) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		res, err := app.UsersGateway.Auth(context.TODO(), &users.Token{
			Val: config.GetState().Username,
		})
		if err != nil {
			return err
		}
		fmt.Println(res.Val)
		return nil
	}
}
