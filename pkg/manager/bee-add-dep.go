package manager

import (
	"errors"
	local "github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
)

func BeeAddDependencies(bee *laruche.Bee, namespaces laruche.Namespaces) error {
	if bee == nil {
		return errors.New("bee == nil")
	}
	ctx := newContext(bee)
	valid := make(laruche.Namespaces, 0)
	_ = namespaces.Map(namespaceValidate(bee, &valid))

	bee.PushDependencies(valid)
	err := local.SaveBee(bee)
	if err != nil {
		return err
	}

	return ctx.dive(bee)
}
