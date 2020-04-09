package manager

import (
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func BeeAddDependencies(bee *laruche.Bee, namespaces laruche.Namespaces) error {
	all := laruche.Append(namespaces, bee.GetSubDependencies()...)
	ctx := newContext(bee)
	fmt.Println(namespaces, bee.Deps)
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
