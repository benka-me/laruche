package absolute

import (
	local "github.com/benka-me/laruche/pkg/get-local"
	remote "github.com/benka-me/laruche/pkg/get-remote"
	"github.com/benka-me/laruche/pkg/laruche"
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

func SaveBee(bee *laruche.Bee) error {
	return nil
}
