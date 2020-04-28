package manager

import (
	"errors"
	absolute "github.com/benka-me/laruche/go-pkg/get-absolute"
	"github.com/benka-me/laruche/go-pkg/git"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func HiveAddDependencies(hive *laruche.Hive, request laruche.Namespaces) error {
	if hive == nil {
		return errors.New("hive == nil")
	}
	valid := make(laruche.Beez, 0)

	err := request.Map(func(i int, req laruche.Namespace) error {
		ok, err := absolute.GetBee(req)
		if err != nil {
			return err
		}

		valid.Push(ok)
		git.Clone(ok.Repo)
		return nil
	})
	if err != nil {
		return err
	}

	err = hive.GetDependencies().Map(func(i int, req laruche.Namespace) error {
		ok, err := absolute.GetBee(req)
		if err != nil {
			return errors.New("hive.yaml: " + err.Error())
		}

		valid.Push(ok)
		return nil
	})
	if err != nil {
		return err
	}

	ctx := newContext(hive)
	return valid.Map(func(i int, toAdd *laruche.Bee) error {
		err = ctx.addDependencyToConsumerAndSave(toAdd)
		if err != nil {
			return err
		}

		err = ctx.dive(toAdd)
		if err != nil {
			return err
		}
		return nil
	})
}
