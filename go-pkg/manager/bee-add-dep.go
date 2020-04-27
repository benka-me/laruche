package manager

import (
	"errors"
	"github.com/benka-me/laruche/go-pkg/generator"
	local "github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/benka-me/laruche/go-pkg/git"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func BeeAddDependencies(bee *laruche.Bee, request laruche.Namespaces) error {
	if bee == nil {
		return errors.New("bee == nil")
	}
	ctx := newContext(bee)
	valid := make(laruche.Beez, 0)
	_ = request.Map(namespaceValidate(bee, &valid))

	bee.PushDependencies(valid.GetDependencies())

	err := generator.GenerateClients(bee)
	if err != nil {
		return err
	}

	_, _ = git.Clone(bee.Repo)

	err = local.SaveBee(bee)
	if err != nil {
		return err
	}

	return ctx.dive(bee)
}
