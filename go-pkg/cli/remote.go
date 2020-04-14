package cli

import (
	"context"
	"fmt"
	local "github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"github.com/benka-me/users/go-pkg/users"
	"github.com/urfave/cli"
)

func push(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		beeOrHive, err := GetOneOfInCurrentDir()
		if err != nil {
			return err
		}

		err = beeOrHive.push(app)
		if err != nil {
			return err
		}

		return nil
	}
}
func publish(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		beeOrHive, err := GetOneOfInCurrentDir()
		if err != nil {
			return err
		}

		err = beeOrHive.publish(app)
		if err != nil {
			return err
		}

		return nil
	}
}
func privatize(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		beeOrHive, err := GetOneOfInCurrentDir()
		if err != nil {
			return err
		}

		err = beeOrHive.privatize(app)
		if err != nil {
			return err
		}

		return nil
	}
}

func (bee Bee) push(app *App) error {
	b := laruche.Bee(bee)
	_, err := app.LarsrvGateway.SetBee(context.TODO(), &laruche.PushBee{
		Bee:   &b,
		Token: &users.Token{Val: app.State.AuthToken},
	})
	if err != nil {
		return err
	}
	fmt.Println("successfuly updated")
	return nil
}

func (bee Bee) publish(app *App) error {
	bee.Public = true
	b := laruche.Bee(bee)
	_ = local.SaveBee(&b)
	_, err := app.LarsrvGateway.PublishBee(context.TODO(), &laruche.PushBee{
		Bee:   &b,
		Token: &users.Token{Val: app.State.AuthToken},
	})
	if err != nil {
		return err
	}
	fmt.Println("successfuly published")
	return nil
}

func (bee Bee) privatize(app *App) error {
	bee.Public = false
	b := laruche.Bee(bee)
	_ = local.SaveBee(&b)
	_, err := app.LarsrvGateway.PrivatizeBee(context.TODO(), &laruche.PushBee{
		Bee:   &b,
		Token: &users.Token{Val: app.State.AuthToken},
	})
	if err != nil {
		return err
	}
	fmt.Println("successfuly privatized")
	return nil
}

func (hive Hive) push(app *App) error {
	h := laruche.Hive(hive)
	_, err := app.LarsrvGateway.SetHive(context.TODO(), &laruche.PushHive{
		Hive:  &h,
		Token: &users.Token{Val: app.State.AuthToken},
	})
	if err != nil {
		return err
	}
	fmt.Println("successfuly updated")
	return nil
}

func (hive Hive) publish(app *App) error {
	hive.Public = true
	h := laruche.Hive(hive)
	_ = local.SaveHive(&h)
	_, err := app.LarsrvGateway.SetHive(context.TODO(), &laruche.PushHive{
		Hive:  &h,
		Token: &users.Token{Val: app.State.AuthToken},
	})
	if err != nil {
		return err
	}
	fmt.Println("successfuly published")
	return nil
}

func (hive Hive) privatize(app *App) error {
	hive.Public = false
	h := laruche.Hive(hive)
	_ = local.SaveHive(&h)
	_, err := app.LarsrvGateway.SetHive(context.TODO(), &laruche.PushHive{
		Hive:  &h,
		Token: &users.Token{Val: app.State.AuthToken},
	})
	if err != nil {
		return err
	}
	fmt.Println("successfuly privatized")
	return nil
}
