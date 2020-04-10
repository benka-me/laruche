package manager

import (
	"errors"
	"github.com/benka-me/laruche/pkg/generator"
	local "github.com/benka-me/laruche/pkg/get-local"
	"github.com/benka-me/laruche/pkg/laruche"
)

func BeeAddDependencies(bee *laruche.Bee, request laruche.Namespaces) error {
	if bee == nil {
		return errors.New("bee == nil")
	}
	ctx := newContext(bee)
	valid := make(laruche.Beez, 0)
	_ = request.Map(namespaceValidate(bee, &valid))

	bee.PushDependencies(valid.GetDependencies())

	err := generator.Clients(bee)
	if err != nil {
		return err
	}

	err = local.SaveBee(bee)
	if err != nil {
		return err
	}

	return ctx.dive(bee)
}
