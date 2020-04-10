package manager

import (
	"errors"
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func HiveAddDependencies(hive *laruche.Hive, namespaces laruche.Namespaces) error {
	if hive == nil {
		return errors.New("hive == nil")
	}
	// TODO: protect from invalid namespaces

	all := laruche.AppendUnique(namespaces, hive.GetDependencies()...)
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

		fmt.Println("will dive: ", toAdd.GetNamespace(), nspace)
		err = ctx.dive(toAdd)
		if err != nil {
			return err
		}
		return nil
	})
}
