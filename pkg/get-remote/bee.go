package remote

import (
	"errors"
	"github.com/benka-me/laruche/pkg/laruche"
)

func GetBee(namespace laruche.Namespace) (*laruche.Bee, error) {
	bee := laruche.Bee{}

	return &bee, errors.New("invalid namespace: " + namespace.String())
}

func SaveBee(bee *laruche.Bee) error {
	return nil
}
