package manager

import (
	"errors"
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func HiveAddDependencies(hive *laruche.Hive, request laruche.Namespaces) error {
	if hive == nil {
		return errors.New("hive == nil")
	}
	valid := make(laruche.Beez, 0)
	_ = request.Map(func(i int, req laruche.Namespace) error {
		ok, err := absolute.GetBee(req)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		valid.Push(ok)
		return nil
	})

	all := laruche.AppendUnique(request, hive.GetDependencies()...)
	ctx := newContext(hive)
	return all.Map(func(i int, nspace laruche.Namespace) error {
		toAdd, err := absolute.GetBee(nspace)
		if err != nil {
			return err
		}

		err = ctx.AddDependencyToConsumerAndSave(toAdd)
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
