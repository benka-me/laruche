package remote

import (
	"github.com/benka-me/laruche/pkg/laruche"
)

func GetBee(namespace laruche.Namespace) (*laruche.Bee, error) {
	bee := laruche.Bee{}
	return &bee, nil
}

func SaveBee(bee *laruche.Bee) error {
	return nil
}
