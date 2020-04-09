package absolute

import (
	local "github.com/benka-me/laruche/pkg/get-local"
	remote "github.com/benka-me/laruche/pkg/get-remote"
	"github.com/benka-me/laruche/pkg/laruche"
)

func GetHive(namespace laruche.Namespace) (*laruche.Hive, error) {
	ret, err := local.GetHive(namespace)
	if err != nil {
		// get remote bee if he's not installed on local machine
		ret, err = remote.GetHive(namespace)
		if err != nil {
			return ret, err
		}
	}
	return ret, nil
}

func SaveHive(hive *laruche.Hive) error {
	return nil
}
