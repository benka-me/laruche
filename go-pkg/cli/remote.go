package cli

import (
	"github.com/urfave/cli"
)

func push(app *App) cli.ActionFunc {
	return func(context *cli.Context) error {
		beeOrHive, err := GetOneOfInCurrentDir()
		if err != nil {
			return err
		}

		err = beeOrHive.push()
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

		err = beeOrHive.publish()
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

		err = beeOrHive.privatize()
		if err != nil {
			return err
		}

		return nil
	}
}

func (bee Bee) push() error {
	panic("implement me")
}

func (bee Bee) publish() error {
	panic("implement me")
}

func (bee Bee) privatize() error {
	panic("implement me")
}

func (hive Hive) push() error {
	panic("implement me")
}

func (hive Hive) publish() error {
	panic("implement me")
}

func (hive Hive) privatize() error {
	panic("implement me")
}
