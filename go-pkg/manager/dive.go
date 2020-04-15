package manager

import (
	"errors"
	"fmt"
	absolute "github.com/benka-me/laruche/go-pkg/get-absolute"
	local "github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func (ctx Context) dive(bee *laruche.Bee) error {
	// check c.Traversed contains child namespace
	if ctx.Traversed.Contains(bee.GetNamespace()) {
		return errors.New(fmt.Sprintf("cycle detected: %v\n%s", ctx.Traversed, bee.GetNamespace()))
	}

	// add current node to traversed
	ctx.Traversed.PushUnique(bee.GetNamespace())

	// concat child to ctx.Consumers
	err := ctx.addDependencyToConsumerAndSave(bee)
	if err != nil {
		return err
	}

	// concat c.Consumers to child.Consumers
	bee.PushConsumers(ctx.Consumers)
	err = local.SaveBee(bee)
	if err != nil {
		return err
	}

	// dive into dependencies
	return bee.MapDependencies(func(i int, nspace laruche.Namespace) error {
		dep, err := absolute.GetBee(nspace)
		if err != nil {
			return err
		}

		err = ctx.dive(dep)
		if err != nil {
			return err
		}
		return nil
	})
}
