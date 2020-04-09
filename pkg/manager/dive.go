package manager

import (
	"fmt"
	local "github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
)

func (ctx Context) install(parent *laruche.Bee, child *laruche.Bee) error {
	var err error
	fmt.Println(ctx)
	fmt.Println(parent.Deps)
	fmt.Println(child.Cons)
	// check c.Traversed contains child namespace

	// concat c.Consumers to child.Consumers

	// add child namespace to traversed

	// add child to parents.Deps

	parent.PushDependency(child.GetNamespace())

	err = local.SaveBee(parent)
	if err != nil {
		return err
	}

	err = local.SaveBee(child)
	if err != nil {
		return err
	}

	fmt.Println(ctx)
	fmt.Println(parent.Deps)
	fmt.Println(child.Cons)
	return nil
}
