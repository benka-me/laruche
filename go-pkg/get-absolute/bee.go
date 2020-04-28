package absolute

import (
	local "github.com/benka-me/laruche/go-pkg/get-local"
	remote "github.com/benka-me/laruche/go-pkg/get-remote"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func GetBee(namespace laruche.Namespace) (*laruche.Bee, error) {
	ret, err := local.GetBee(namespace)
	if err != nil {
		// get remote bee if he's not installed on local machine
		ret, err = remote.GetBee(namespace)
		if err != nil {
			return ret, err
		}
	}
	return ret, nil
}

func GetBeez(namespaces laruche.Namespaces) (*laruche.Beez, error) {
	ret := make(laruche.Beez, 0)

	err := namespaces.Map(func(i int, namespace laruche.Namespace) error {
		one, err := GetBee(namespace)
		if err != nil {
			return err
		}
		ret = append(ret, one)
		return nil
	})
	return &ret, err
}
