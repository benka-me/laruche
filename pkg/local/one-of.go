package local

import (
	"errors"
	"github.com/benka-me/laruche/pkg/oneof"
)

func GetOneOfCurrentDir() (oneof.OneOf, error) {
	hive, errHive := GetHiveCurrentDir()
	bee, errBee := GetBeeCurrentDir()

	if errHive != nil && errBee != nil {
		return nil, errors.New("neither bee.yaml or hive.yaml found on this folders")
	} else if errHive == nil && errBee == nil {
		return nil, errors.New("both bee.yaml or hive.yaml found on this folders")
	} else if errHive == nil {
		return oneof.Hive(*hive), nil
	} else if errBee == nil {
		return oneof.Bee(*bee), nil
	}
	return nil, nil
}
