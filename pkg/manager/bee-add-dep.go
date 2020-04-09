package manager

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/local"
)

func BeeAddDependencies(bee *laruche.Bee, namespaces laruche.Namespaces) error {
	ctx := Context{}
	namespaces.Map(func(i int, nspace laruche.Namespace) {
		toAdd, err := local.GetBee(string(nspace))
		if err != nil {
			return
		}
		ctx.Dive(bee, toAdd)
	})
	return nil
}
