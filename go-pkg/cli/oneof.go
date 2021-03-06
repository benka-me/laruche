package cli

import (
	"errors"
	"github.com/benka-me/laruche/go-pkg/get-local"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

type OneOf interface {
	AddDep(bool, laruche.Namespaces) error
	push(app *App) error
	publish(app *App) error
	privatize(app *App) error
}

type Bee laruche.Bee
type Hive laruche.Hive

func GetOneOfInCurrentDir() (OneOf, error) {
	hive, errHive := local.GetHiveCurrentDir()
	bee, errBee := local.GetBeeCurrentDir()

	if errHive != nil && errBee != nil {
		return nil, errors.New("neither bee.yaml or hive.yaml found on this folders")
	} else if errHive == nil && errBee == nil {
		return nil, errors.New("both bee.yaml or hive.yaml found on this folders")
	} else if errHive == nil {
		return Hive(*hive), nil
	} else if errBee == nil {
		return Bee(*bee), nil
	}
	return nil, nil
}
