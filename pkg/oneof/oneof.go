package oneof

import (
	"github.com/benka-me/laruche/pkg/laruche"
	"github.com/benka-me/laruche/pkg/manager"
)

type OneOf interface {
	AddDep(bool, laruche.Namespaces) error
}

type Bee laruche.Bee
type Hive laruche.Hive

func (bee Bee) AddDep(depMode bool, namespaces laruche.Namespaces) error {
	if depMode {
		namespaces = laruche.Bee(bee).GetSubDependencies()
	}

	var lb = laruche.Bee(bee)
	return manager.BeeAddDependencies(&lb, namespaces)
}

func (hive Hive) AddDep(depMode bool, namespaces laruche.Namespaces) error {
	if depMode {
		namespaces = laruche.Hive(hive).GetDependencies()
	}

	var lh = laruche.Hive(hive)
	return manager.HiveAddDependencies(&lh, namespaces)
}
