package manager

import (
	"errors"
	"fmt"
	absolute "github.com/benka-me/laruche/pkg/get-absolute"
	local "github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
)

func (ctx Context) recursion(bee *laruche.Bee) error {
	fmt.Println("dive on " + bee.GetNamespace())
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
	err = local.SaveBee(bee)
	if err != nil {
		return err
	}

	// map dependencies and run same recursion
	return bee.MapDependencies(func(i int, nspace laruche.Namespace) error {
		dep, err := absolute.GetBee(nspace)
		if err != nil {
			return err
		}

		err = ctx.recursion(dep)
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
	ctx := newContext(bee)
	// TODO: protect from self-dependency
	// TODO: protect from invalid namespace

	bee.PushDependencies(namespaces)
	err := local.SaveBee(bee)
	if err != nil {
		return err
	}

	return ctx.recursion(bee)
}
