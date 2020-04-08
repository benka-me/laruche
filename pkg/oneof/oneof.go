package oneof

import (
	"fmt"
	"github.com/benka-me/laruche/pkg/laruche"
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

	fmt.Println(namespaces)
	return nil
}

func (hive Hive) AddDep(depMode bool, namespaces laruche.Namespaces) error {
	if depMode {
		namespaces = laruche.Hive(hive).GetDependencies()
	}
	return nil
}
