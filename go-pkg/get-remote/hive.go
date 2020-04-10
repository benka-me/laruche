package remote

import (
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func GetHive(namespace laruche.Namespace) (*laruche.Hive, error) {
	hive := laruche.Hive{}
	return &hive, nil
}

func SaveHive(hive *laruche.Hive) error {
	return nil
}
