package manager

import (
	local "github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
)

type Context struct {
	Traversed laruche.Namespaces
	Consumers laruche.Namespaces
}

func newContext(oneOf laruche.OneOf) Context {
	ret := Context{
		Traversed: make(laruche.Namespaces, 0),
		Consumers: oneOf.GetConsumers(),
	}

	return ret
}

func (ctx *Context) AddDependencyToConsumerAndSave(bee *laruche.Bee) error {
	return ctx.Consumers.Map(func(i int, nspace laruche.Namespace) error {
		h, err := local.GetHive(nspace)
		if err != nil {
			return err
		}
		h.Deps[bee.GetNamespaceStr()] = &laruche.Dep{
			Port:         bee.Port,
			Dev:          "localhost",
			Prod:         "127.0.0.1",
			PkgName:      bee.PkgName,
			PkgNameCamel: bee.PkgNameCamel,
		}

		err = local.SaveHive(h)
		if err != nil {
			return err
		}
		return nil
	})
}
