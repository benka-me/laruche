package manager

import (
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func HiveAddDependencies(hive *laruche.Hive, namespaces laruche.Namespaces) error {
	ctx := newContext(nil)

	return namespaces.Map(func(i int, nspace laruche.Namespace) error {
		toAdd, err := absolute.GetBee(nspace)
		if err != nil {
			return err
		}

		fmt.Println(ctx, toAdd)
		return nil
	})
}
