package manager

import (
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func namespaceValidate(bee *laruche.Bee, valid *laruche.Beez) laruche.NamespaceIter {
	return func(i int, nspace laruche.Namespace) error {
		if nspace == bee.GetNamespace() {
			fmt.Println(nspace + " cannot be dependency of himself")
			return nil
		}
		try, err := absolute.GetBee(nspace)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		valid.Push(try)
		return nil
	}
}
