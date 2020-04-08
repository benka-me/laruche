package oneof

import (
	"errors"
	"github.com/benka-me/laruche/pkg/local"
)

func GetOneOfCurrentDir() (OneOf, error) {
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
