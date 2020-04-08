package manager

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/local"
)

func BeeAddDependencies(bee *laruche.Bee, namespaces laruche.Namespaces) error {
	namespaces.Map(func(i int, nspace laruche.Namespace) {
		try, err := local.GetBee(string(nspace))
		if err != nil {
			return
		}
		fmt.Println(try)
	})
	return nil
}
