package manager

import (
	"errors"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func BeeAddDependencies(bee *laruche.Bee, namespaces laruche.Namespaces) error {
	if bee == nil {
		return errors.New("bee == nil")
	}
	all := laruche.Append(namespaces, bee.GetSubDependencies()...)
	ctx := newContext(bee)
	return all.Map(func(i int, nspace laruche.Namespace) error {
		toAdd, err := absolute.GetBee(nspace)
		if err != nil {
			return err
		}

		err = ctx.install(bee, toAdd)
		if err != nil {
			return err
		}

		return nil
	})
}
