package manager

import (
	"errors"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func HiveAddDependencies(hive *laruche.Hive, namespaces laruche.Namespaces) error {
	if hive == nil {
		return errors.New("hive == nil")
	}

	all := laruche.AppendUnique(namespaces, hive.GetDependencies()...)
	ctx := newContext(hive)
	return all.Map(func(i int, nspace laruche.Namespace) error {
		toAdd, err := absolute.GetBee(nspace)
		if err != nil {
			return err
		}

		err = ctx.AddDependencyToConsumer(toAdd)
		if err != nil {
			return err
		}

		return nil
	})
}
