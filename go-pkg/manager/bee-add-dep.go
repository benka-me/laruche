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

	_ = valid.Map(func(i int, b *laruche.Bee) error {
		_, _ = git.Clone(b.Repo)
		return nil
	})

	err := generator.GenerateClients(bee)
	if err != nil {
		return err
	}

	err = local.SaveBee(bee)
	if err != nil {
		return err
	}

	return ctx.dive(bee)
}
