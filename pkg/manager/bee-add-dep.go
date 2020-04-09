package manager

import (
	"errors"
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	"github.com/benka-me/laruche/pkg/laruche"
)

func (ctx Context) beeAddRecursion(bee *laruche.Bee) error {
	// check c.Traversed contains child namespace
	if ctx.Traversed.Contains(bee.GetNamespace()) {
		return errors.New(fmt.Sprintf("cycle detected: %v\n%s", ctx.Traversed, bee.GetNamespace()))
	}

	// add current node to traversed
	ctx.Traversed.Push(bee.GetNamespace())

	// concat child to ctx.Consumers
	err := ctx.AddDependencyToConsumer(bee)
	if err != nil {
		return err
	}

	// concat c.Consumers to child.Consumers
	bee.PushConsumers(ctx.Consumers)

	// map dependencies and run same recursion
	return bee.MapDependencies(func(i int, nspace laruche.Namespace) error {
		dep, err := absolute.GetBee(nspace)
		if err != nil {
			return err
		}

		err = ctx.beeAddRecursion(dep)
		if err != nil {
			return err
		}
		return nil
	})
}

func BeeAddDependencies(bee *laruche.Bee, namespaces laruche.Namespaces) error {
	if bee == nil {
		return errors.New("bee == nil")
	}
	bee.PushDependencies(namespaces)
	ctx := newContext(bee)

	return ctx.beeAddRecursion(bee)
}
