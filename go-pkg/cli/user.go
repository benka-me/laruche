package cli

import (
	"context"
	"fmt"
	"github.com/benka-me/laruche/go-pkg/cli/scan"
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

		res, err := app.Users.Login(context.TODO(), req)
		if err != nil {
			return err
		}
		fmt.Println("Success")
		app.State.SetAuth(res.Data.Username, res.Auth)
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
			Val: app.State.AuthToken,
		})
		if err != nil {
			return err
		}
		if res.Val {
			fmt.Println("You are authenticated as", app.State.Username)
		}
		return nil
	}
}
